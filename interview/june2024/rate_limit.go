package june2024

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

var (
	SuccessResponse = Resposne{
		Status:  http.StatusOK,
		Message: "OK",
	}

	TooManyRequestResponse = Resposne{
		Status:  http.StatusTooManyRequests,
		Message: "Too many requests",
	}
)

type Resposne struct {
	Status  int32
	Message string
}

const ResponseFormat = "{status: %d, message: %s}"

func (r *Resposne) String() string {
	return fmt.Sprintf(ResponseFormat, r.Status, r.Message)
}

type RateLimiter struct {
	// Config
	window int64 // in seconds
	limit  int64

	// Storage
	data sync.Map // Key -> *Dequeue
}

type DoubleLinkedListNode[T any] struct {
	Value      T
	Next, Prev *DoubleLinkedListNode[T]
}

func NewDequeue[T any]() *Dequeue[T] {
	return &Dequeue[T]{}
}

type Dequeue[T any] struct {
	length     int
	head, tail *DoubleLinkedListNode[T]
}

func (d *Dequeue[T]) Len() int {
	return d.length
}

func (d *Dequeue[T]) Front() (v T) {
	return d.head.Value
}

func (d *Dequeue[T]) PopFront() {
	if d.Len() == 1 {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.Next
		d.head.Prev = nil
	}
	d.length--
}

func (d *Dequeue[T]) PushBack(v T) {
	newNode := &DoubleLinkedListNode[T]{
		Value: v,
		Prev:  d.tail,
	}
	if d.Len() == 0 {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.Next = newNode
		d.tail = newNode
	}
	d.length++
}

// Help to check if we allow the request key at the given time
// Will only log the successful request
func (r *RateLimiter) Allow(_ context.Context, key string, at int64) bool {
	val, ok := r.data.Load(key)
	if !ok {
		return true
	}
	dequeue, _ := val.(*Dequeue[int64])
	for dequeue.Len() > 0 && dequeue.Front() <= at-r.window {
		dequeue.PopFront()
	}
	return dequeue.Len() < int(r.limit)
}

func (r *RateLimiter) Record(_ context.Context, key string, at int64) {
	val, ok := r.data.Load(key)
	if !ok {
		newDequeue := NewDequeue[int64]()
		newDequeue.PushBack(at)
		r.data.Store(key, newDequeue)
		return
	}
	dequeue, _ := val.(*Dequeue[int64])
	dequeue.PushBack(at)
}

func getRequestStatus(requests []string) []string {
	rateLimiter1 := RateLimiter{
		window: 5,
		limit:  2,
	}
	rateLimiter2 := RateLimiter{
		window: 30,
		limit:  5,
	}
	responses := make([]string, len(requests))
	ctx := context.Background()
	for i := range requests {
		if !rateLimiter1.Allow(ctx, requests[i], int64(i)) {
			responses[i] = TooManyRequestResponse.String()
			continue
		}
		if !rateLimiter2.Allow(ctx, requests[i], int64(i)) {
			responses[i] = TooManyRequestResponse.String()
			continue
		}
		responses[i] = SuccessResponse.String()
		rateLimiter1.Record(ctx, requests[i], int64(i))
		rateLimiter2.Record(ctx, requests[i], int64(i))
	}
	return responses
}
