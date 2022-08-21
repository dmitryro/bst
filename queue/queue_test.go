package queue

import (
    "fmt"
    "testing"
)

// Test the dequeueing from queue 
// Place one test value and remove it 
// Assert success if queue is empty
// Assert failure if queue is not empty
func TestDequeue(t *testing.T) {
   q := new(Queue[*int])
   v := 1
   q.Enqueue(&v)
   q.Dequeue()
   length := q.Size()

   if length != 0 {
       t.Errorf("After dequeue Queue size is incorrect, got: %d, expected: %d.", length, 0)
   } else {
       fmt.Printf("Dequeue size - %d\n", length)
   }
}

// Test the enqueueing into  queue 
// Place one test value 
// Assert success if queue length is 1
// Assert failure if queue length is not 1
func TestEnqueue(t *testing.T) {
   q := new(Queue[*int])
   v := 1
   q.Enqueue(&v)
   length := q.Size()

   if length != 1 {
       t.Errorf("After enqueueing into Queue size is incorrect, got: %d, expected: %d.", length, 1)
   } else {
       fmt.Printf("Enqueue size - %d\n", length)
   }
}

// Test the queue sizing
// Add values 1, 2, 3 to the queue
// Assert success if queue size is 3
// Assert failure if queue size is not 3
func TestQueueSize(t *testing.T) {
   q := new(Queue[*int])
   a := 1
   b := 2
   c := 3
   q.Enqueue(&a)
   q.Enqueue(&b)
   q.Enqueue(&c)
   length := q.Size()

   if length != 3 {
       t.Errorf("After enqueueing into Queue size is incorrect, got: %d, expected: %d.", length, 3)
   } else {
       fmt.Printf("Queue Size - %d\n", length)
   }
}
