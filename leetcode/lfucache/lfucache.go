package lfucache

type LFUCache struct {
	list  *lfuDList
	cap   int
	count int
	items map[int]*storeItem
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		list:  &lfuDList{},
		cap:   capacity,
		items: make(map[int]*storeItem),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	// find if present
	item, ok := this.items[key]
	if ok {
		newHead := item.touch()
		if newHead != nil {
			this.list.head = newHead
		}
		return item.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	// find if present
	item, ok := this.items[key]
	if ok {
		// present: update the value
		item.value = value
		// increase the count
		newHead := item.touch()
		if newHead != nil {
			this.list.head = newHead
		}
		return
	}
	// not exist
	// check capacity
	if this.count == this.cap {
		// remove the least used key
		evictedItem := this.list.evict()
		this.count--
		delete(this.items, evictedItem.key)
	}

	this.count++
	this.items[key] = this.list.addItem(key, value)
}

type lfuDList struct {
	head *lfuDNode
}

// Assume:
// * l != nil
// * l.head != nil
func (l *lfuDList) evict() *storeItem {
	if l.head.headItem == l.head.tailItem {
		// remove the head
		item := l.head.headItem
		if l.head.next == nil {
			l.head.selfRemove()
		} else {
			newHead := l.head.next
			l.head.selfRemove()
			l.head = newHead
		}
		return item
	}
	// remove the first key in head
	item := l.head.popHeadItem()
	return item
}

// Assume:
// * l != nil
func (l *lfuDList) addItem(key, value int) *storeItem {
	item := &storeItem{
		key:   key,
		value: value,
	}
	if l.head == nil {
		// create new node
		newNode := &lfuDNode{
			freq:     1,
			headItem: item,
			tailItem: item,
		}
		item.pNode = newNode
		l.head = newNode
		return item
	}
	if l.head.freq == 1 {
		l.head.addItem(item)
		return item
	}
	// create new node
	newNode := &lfuDNode{
		freq:     1,
		headItem: item,
		tailItem: item,
	}
	item.pNode = newNode
	l.head.insertBefore(newNode)
	l.head = newNode
	return item
}

type lfuDNode struct {
	freq     int
	headItem *storeItem
	tailItem *storeItem
	prev     *lfuDNode
	next     *lfuDNode
}

// Assume:
// * n != nil
// * newNode != nil
func (n *lfuDNode) insertBefore(newNode *lfuDNode) {
	newNode.next = n
	newNode.prev = n.prev
	n.prev = newNode
}

// Assume:
// * n != nil
// * newNode != nil
func (n *lfuDNode) insertAfter(newNode *lfuDNode) {
	newNode.prev = n
	newNode.next = n.next
	n.next = newNode
}

// Assume:
// * n != nil
func (n *lfuDNode) selfRemove() {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.prev = nil
	n.next = nil
}

// Assume:
// * n != nil
func (n *lfuDNode) addItem(item *storeItem) *lfuDNode {
	item.pNode = n
	n.appendStoreItem(item)
	return n
}

// Assume:
// * n != nil
func (n *lfuDNode) appendStoreItem(i *storeItem) {
	i.pNode = n
	i.prev = n.tailItem
	n.tailItem.next = i
	n.tailItem = i
}

// Assume:
// * n != nil
// * n.headItem != nil
// * n.tailItem != nil
// * n.headItem != n.tailItem
func (n *lfuDNode) popHeadItem() *storeItem {
	item := n.headItem
	next := item.next
	item.selfRemove()
	n.headItem = next
	return item
}

type storeItem struct {
	key   int
	value int
	pNode *lfuDNode
	prev  *storeItem
	next  *storeItem
}

// Assume:
// * i != nil
// * i.pNode != nil
func (i *storeItem) selfRemove() {
	if i.prev != nil {
		i.prev.next = i.next
	} else if i.next != nil {
		i.pNode.headItem = i.next
	}
	if i.next != nil {
		i.next.prev = i.prev
	} else if i.prev != nil {
		i.pNode.tailItem = i.prev
	}
	i.prev = nil
	i.next = nil
}

func (i *storeItem) touch() *lfuDNode {
	n := i.pNode
	// Only 1 item in the node
	if n.headItem == n.tailItem {
		// no next node or next node is far
		if n.next == nil || n.next.freq != n.freq+1 {
			n.freq++
			return nil
		}
		// move the item to next node
		next := n.next
		i.pNode = next
		next.appendStoreItem(i)
		n.selfRemove()
		// need to update the head of list
		if next.prev == nil {
			return next
		}
		return nil
	}
	// Many items in the node
	if n.next == nil || n.next.freq != n.freq+1 {
		// remove the key here
		i.selfRemove()
		// create new node
		newNode := lfuDNode{
			freq:     n.freq + 1,
			headItem: i,
			tailItem: i,
		}
		i.pNode = &newNode
		n.insertAfter(&newNode)

		return nil
	}
	// remove the key here
	i.selfRemove()
	// move the item to next node
	i.pNode = n.next
	n.next.appendStoreItem(i)
	return nil
}
