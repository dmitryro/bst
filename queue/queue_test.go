package queue

import (
    "fmt"
    "testing"
)

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
