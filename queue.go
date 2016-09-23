package Queue

import (
	"errors"
	"fmt"
	"reflect"
)

type Queue struct {
	list []interface{}
	MAX  int
}

// returns a pointer to Queue
func NewQueue(MaxSize int) *Queue {
	q := new(Queue)
	q.list = make([]interface{}, 0)
	q.MAX = MaxSize
	return q
}

// returns length of queue
func (q *Queue) Len() int {
	return len(q.list)
}

// appends an item to the rear of queue
func (q *Queue) Enqueue(item interface{}) (err error) {
	if q.Len() == q.MAX {
		err = errors.New(fmt.Sprintf("Buffer Overflow: MaxSize Allowed: %d", q.MAX))
	} else {
		q.list = append(q.list, item)
	}
	return
}

// return the item deleted from the queue
func (q *Queue) Dequeue() (item interface{}, err error) {
	if q.Len() == 0 {
		err = errors.New(fmt.Sprintf("Buffer Underflow"))
	} else {
		item = q.list[0]
		q.list = q.list[:0+copy(q.list[0:], q.list[1:])]
	}
	return
}

// returns nil if queue is empty else returns the queue as a slice
// panics in case if items are of different types in queue.
func (q *Queue) ToSlice() interface{} {
	if q.Len() == 0 {
		return nil
	}
	Type := reflect.TypeOf(q.list[0])
	list := reflect.MakeSlice(reflect.SliceOf(Type), 0, q.Len())
	for _, item := range q.list {
		list = reflect.Append(list, reflect.ValueOf(item))
	}
	return list.Interface()
}
