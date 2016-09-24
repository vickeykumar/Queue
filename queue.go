package Queue

import (
	"fmt"
	"reflect"
	"sync"
)

type Queue struct {
	list []interface{}
	mu   sync.RWMutex
}

// returns a pointer to threadSafe Queue
func NewQueue(MaxSize int) *Queue {
	q := new(Queue)
	q.list = make([]interface{}, 0, MaxSize)
	return q
}

// returns length of queue
func (q *Queue) Len() int {
	return len(q.list)
}

// returns Max capacity of Queue
func (q *Queue) Cap() int {
	return cap(q.list)
}

// Clears the Queue
func (q *Queue) Clear() {
	capacity := cap(q.list)
	q.mu.Lock()
	q.list = nil
	q.list = make([]interface{}, 0, capacity)
	q.mu.Unlock()
}

// appends an item to the rear of queue
func (q *Queue) Enqueue(item interface{}) (err error) {
	if q.Len() == q.Cap() {
		err = fmt.Errorf("Buffer Overflow: MaxSize Allowed: %d", q.Cap())
	} else {
		q.mu.Lock()
		q.list = append(q.list, item)
		q.mu.Unlock()
	}
	return
}

// returns the item deleted from the queue and error if any.
func (q *Queue) Dequeue() (item interface{}, err error) {
	if q.Len() == 0 {
		err = fmt.Errorf("Buffer Underflow")
	} else {
		q.mu.Lock()
		item = q.list[0]
		q.list = q.list[:0+copy(q.list[0:], q.list[1:])]
		q.mu.Unlock()
	}
	return
}

// returns true if Queue contains the item
func (q *Queue) Contains(item interface{}) (flag bool) {
	q.mu.RLock()
	for _, i := range q.list {
		if i == item {
			flag = true
			break
		}
	}
	q.mu.RUnlock()
	return
}

// returns nil if queue is empty else returns the queue as a slice
// error in case if items are of different types in queue.
func (q *Queue) ToSlice() (slice interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			slice = q.list
			err = fmt.Errorf("%v", r)
			return
		}
	}()
	if q.Len() == 0 {
		return nil, err
	}
	q.mu.RLock()
	templist := q.list
	q.mu.RUnlock()
	Type := reflect.TypeOf(q.list[0])
	list := reflect.MakeSlice(reflect.SliceOf(Type), 0, q.Len())
	for _, item := range templist {
		list = reflect.Append(list, reflect.ValueOf(item))
	}
	return list.Interface(), err
}
