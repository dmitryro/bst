package queue

type Queue[T any] []T


func (q *Queue[T]) Enqueue(v T) {
    /* Enqueue a new item to the queue */
	*q = append(*q, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    /* Dequeue an item from the queue */
    size := len(*q)
	if size == 0 {
		var v T
		return v, false
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v, true
}

func (q *Queue[T]) Size() int {
    return len(*q)
}
