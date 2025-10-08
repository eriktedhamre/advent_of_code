package utils

type Queue[T any] struct {
	data []T
	head int
}

func (q *Queue[T]) Enqueue(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head >= len(q.data) {
		var zero T
		return zero, false
	}
	v := q.data[q.head]
	q.head++

	// Optional: reclaim memory periodically
	if q.head > 64 && q.head*2 >= len(q.data) {
		q.data = append([]T(nil), q.data[q.head:]...)
		q.head = 0
	}

	return v, true
}

func (q *Queue[T]) Len() int {
	return len(q.data) - q.head
}
