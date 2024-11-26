package slice

type Queue[T any] struct {
	queue []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		queue: make([]T, 0, 10),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.queue = append(q.queue, item)
}

func (q *Queue[T]) Dequeue() (zero T) {
	if len(q.queue) == 0 {
		return zero
	}
	val := q.queue[0]
	q.queue = q.queue[1:]
	return val
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.queue) == 0
}
