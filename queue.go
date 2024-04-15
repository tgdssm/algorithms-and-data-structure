package main

type Queue struct {
	queue []interface{}
	start int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Size() int {
	return len(q.queue)
}

func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}

func (q *Queue) Enqueue(content interface{}) {
	q.queue = append(q.queue, content)
}

func (q *Queue) Dequeue() (removed interface{}) {
	if !q.isEmpty() {
		removed = q.queue[0]
		q.start++
		if q.start > q.Size()/2 {
			newQueue := make([]interface{}, len(q.queue)-q.start)
			copy(newQueue, q.queue[q.start:])
			q.queue = newQueue
			q.start = 0
		}
		return
	}
	return nil
}

func (q *Queue) Peek() interface{} {
	if !q.isEmpty() {
		return q.queue[q.start]
	}
	return nil
}

func (q *Queue) Clear() {
	q.queue = nil
	q.start = 0
}
