package queue

// Generic Queue data type to be used for BFS and other searches or store
type Queue[T any] []T


// Enqueue a new item to the queue
func (q *Queue[T]) Enqueue(v T) {
    /* Enqueue a new item to the queue */
	*q = append(*q, v)
}

// Dequeue an item from the top of the queue
// Return false if size is 0, true if successfully dequeued
// Return the dequeued item
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

// Return the current size of the queue
func (q *Queue[T]) Size() int {
    return len(*q)
}
