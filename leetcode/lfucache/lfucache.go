package lfucache

type LFUCache struct {
	list  *lfuDList
	cap   int
	count int
	m     map[int]*LFUDNode
	cache map[int]int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		list:  &lfuDList{},
		cap:   capacity,
		m:     make(map[int]*LFUDNode),
		cache: make(map[int]int),
	}
}

func (this *LFUCache) Get(key int) int {
	if this.cap == 0 {
		return -1
	}
	// find if present
	node, ok := this.m[key]
	if ok {
		newNode, needUpdateHead := node.touchKey(key)
		if needUpdateHead {
			this.list.head = newNode
		}
		this.m[key] = newNode
		return this.cache[key]
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	// find if present
	node, ok := this.m[key]
	if ok {
		// present: update the value
		this.cache[key] = value
		// increase the count
		newNode, needUpdateHead := node.touchKey(key)
		if needUpdateHead {
			this.list.head = newNode
		}
		this.m[key] = newNode
	} else {
		// check capacity
		if this.count == this.cap {
			// remove the least used key
			evictedKey := this.list.evict()
			this.count--
			delete(this.cache, evictedKey)
			delete(this.m, evictedKey)
		}

		this.count++
		newNode := this.list.addKey(key)
		this.m[key] = newNode
		this.cache[key] = value
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type lfuDList struct {
	head *LFUDNode
}

func (l *lfuDList) evict() int {
	if l == nil || l.head == nil {
		panic("invalid list")
	}
	if len(l.head.keys) == 1 {
		// remove the head
		key := l.head.keys[0]
		if l.head.next == nil {
			l.head.selfRemove()
		} else {
			newHead := l.head.next
			l.head.selfRemove()
			l.head = newHead
		}
		return key
	}
	// remove the first key in head
	key := l.head.keys[0]
	l.head.keys = l.head.keys[1:]
	return key
}

func (l *lfuDList) addKey(key int) *LFUDNode {
	if l == nil {
		panic("invalid list")
	}
	if l.head == nil {
		// create new node
		newNode := LFUDNode{
			freq: 1,
			keys: []int{key},
		}
		l.head = &newNode
		return &newNode
	}
	if l.head.freq == 1 {
		return l.head.addKey(key)
	}
	// create new node
	newNode := LFUDNode{
		freq: 1,
		keys: []int{key},
	}
	l.head.insertBefore(&newNode)
	l.head = &newNode
	return &newNode
}

type LFUDNode struct {
	freq int
	keys []int // LRU queue (FIFO)
	prev *LFUDNode
	next *LFUDNode
}

func (n *LFUDNode) insertBefore(newNode *LFUDNode) {
	if n == nil || newNode == nil {
		panic("invalid node")
	}
	newNode.next = n
	newNode.prev = n.prev
	n.prev = newNode
}

func (n *LFUDNode) insertAfter(newNode *LFUDNode) {
	if n == nil || newNode == nil {
		panic("invalid node")
	}
	newNode.prev = n
	newNode.next = n.next
	n.next = newNode
}

func (n *LFUDNode) selfRemove() {
	if n == nil {
		panic("invalid node")
	}
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

func (n *LFUDNode) addKey(key int) *LFUDNode {
	if n == nil {
		panic("invalid node")
	}
	n.keys = append(n.keys, key)
	return n
}

func (n *LFUDNode) touchKey(key int) (newNode *LFUDNode, needUpdateHead bool) {
	if n == nil || len(n.keys) == 0 {
		panic("invalid node")
	}
	if len(n.keys) == 1 {
		if n.next == nil || n.next.freq != n.freq+1 {
			n.freq++
			return n, false
		} else {
			// move the item to next node
			// need to update the head of list
			next := n.next
			next.keys = append(next.keys, key)
			n.selfRemove()
			return next, next.prev == nil
		}
	}
	for i := range n.keys {
		if n.keys[i] == key {
			if n.next == nil || n.next.freq != n.freq+1 {
				// create new node
				newNode := LFUDNode{
					freq: n.freq + 1,
					keys: []int{key},
				}
				n.insertAfter(&newNode)
				// remove the key here
				n.keys = n.keys[:i+copy(n.keys[i:], n.keys[i+1:])]
				return &newNode, false
			} else {
				// move the item to next node
				n.next.keys = append(n.next.keys, key)
				// remove the key here
				n.keys = n.keys[:i+copy(n.keys[i:], n.keys[i+1:])]
				return n.next, false
			}
		}
	}
	panic("invalid case")
}
