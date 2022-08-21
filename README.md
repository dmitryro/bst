# Go Implementation of Binary Search Tree
Binary Tree Implementation using Go, Queue and Tree data structures, Breadth First Search for tree traversal.

## Getting Started


## Versions used


## Installing Go 1.18 or later


## Initial Assumptions
To successfully run the project some initial assumptions need to be made:

### **Valid Go Version**
**Go 1.18 SDK** or later must be installed on the system.

### **GOROOT** 
    should point to the proper Go SDK location.

### **GOPATH** 
should be pointing to  a valid workspace.

### **Queue** 
data structure is used to implement **BFS** search. We use a custom implementation, while alternative 
      third party implementations exist.

### **Tree** and **Queue** 
behaviors are covered by unit tests, to run which please follow the instructions below.
     

## How to Import
### Importing the Queue
#### If you're using the queue in your go file, include the following import:
```
import (
      "dmitryro.com/queue"
)
```
Please see example in [tree/tree.go](https://github.com/dmitryro/gotests/blob/master/tree/tree.go) file.

### Importing the Binary Search Tree
#### If you're using the queue in your go file, include the following import:
```
import (
      "dmitryro.com/tree"
)
```
Please see example in [main.go](https://github.com/dmitryro/gotests/blob/master/main.go) file.





## [Queue](https://github.com/dmitryro/gotests/blob/master/queue/queue.go)
### Data Structures:
#### Generic Queue Type is used to allow reuse for different types
```
type Queue[T any] []T
```
* Please note - versions before 1.18 do not support generic types, so please make sure at least Go version 1.18 is used.


### Methods
#### [**Enqueue**](https://github.com/dmitryro/gotests/blob/master/queue/queue.go#L6):
  - **Description**: Add a new item to the **Queue**.

  - **Signature**: ``` func (q *Queue[T]) Enqueue(v T)  ```

  - **Returned Value**: Returns no values


#### [**Dequeue**](https://github.com/dmitryro/gotests/blob/master/queue/queue.go#L11): 
  - **Description**: Remove an item from the top of the **Queue**.

  - **Signature**: ``` func (q *Queue[T]) Dequeue() (T, bool)  ```

  - **Returned Value**: Returns **true** or **false** in case of **success** or **falure** to dequeue, and the top item removed from teh queue.



#### [**Size**](https://github.com/dmitryro/gotests/blob/master/queue/queue.go#L23)
  - **Description**: Return the current size of the **Queue**.
  - **Signature**: ``` func (q *Queue[T]) Size() int ```
  - **Returned Value**: Returns **int** value of the queue size.


### Tests
#### [**TestDequeue**](https://github.com/dmitryro/gotests/blob/master/queue/queue_test.go#L8):
  - **Description**: Verify the item is dequeued from the **Queue**.

  - **Signature**: ``` func TestDequeue(t *testing.T)  ```

  - **Assumptions**:  A new queue has been initialized, a value has been enqueued and then dequeued from the queue.

  - **Asserted Success**: The size of queue after value is removed *is* 0.

  - **Asserted Failure**: The size of queue after value is removed is *not* 0. 


#### [**TestEnqueue**](https://github.com/dmitryro/gotests/blob/master/queue/queue_test.go#L22):
  - **Description**: Verify the item is enqueued to the **Queue**.

  - **Signature**: ``` func TestEnqueue(t *testing.T)  ```

  - **Assumptions**: A new queue has been initialized and a value is enqueued into the queue.

  - **Asserted Success**: The size of queue *is* 1.

  - **Asserted Failure**: The size of queue is *not* 1.


#### [**TestQueueSize**](https://github.com/dmitryro/gotests/blob/master/queue/queue_test.go#L35):
  - **Description**: Verify the size of the **Queue** is as expected.

  - **Signature**: ``` func TestQueueSize(t *testing.T)  ```

  - **Assumptions**: For the sake of unit test, a new queue is initialized and 3 values added to the queue

  - **Asserted Success**: The size of queue *is* 3.

  - **Asserted Failure**: The size of queue is *not* 3.




## [Binary Search Tree](https://github.com/dmitryro/gotests/blob/master/tree/tree.go)
### Data Structures
#### TreeNode

This type is used to represent a tree node with value, left and right children, and depth:
```
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
    Depth int
}
``` 

#### NodeItem

This type is used to return the value and the depth of a found node or list of nodes.
```
type NodeItem struct {
    Val int
    Depth int
}
```

#### BSTree

This type represents the tree itself - pointer to the root node and the current size of the tree. Using it
allows better design of tree-specific methods and behaviors.
```
type BSTree struct {
    Size int
    Root *TreeNode
}
```

### Methods
####


## Tests
###


### Running Tests
#### To run Queue tests 

```
   cd ./queue
```
and execute:
```
   go test
```

#### To run Tree tests 

```
   cd ./queue
```
and execute:
```
   go test
```

## Author
Dmitry Roiitman 



