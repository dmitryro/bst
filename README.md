# Go Implementation of Binary Search Tree
Binary Tree Implementation using Go, Queue and Tree data structures, Breadth First Search for tree traversal.

## Getting Started


## Versions used


## Installing Go 1.18 or later


## Initial Assumptions


## How to Import







## Queue 
### Data Structures:
#### Generic Queue Type is used to allow reuse for different types
```
type Queue[T any] []T
```
* Please note - versions before 1.18 do not support generic types, so please make sure at least Go version 1.18 is used.


### Methods
#### **Enqueue**
**Description**: Add a new item to the **Queue**.
**Signature**: ``` func (q *Queue[T]) Enqueue(v T)  ```
**Returned Value**: Returns no values


#### **Dequeue** 
#### Description: Remove an item from the top of the **Queue**.
#### Signature: ``` func (q *Queue[T]) Dequeue() (T, bool) ```
#### Returns **true** or **false** in case of **success** or **falure** to dequeue, and the top item removed from teh queue.


#### **Size**
#### Description: Return the current size of the **Queue**.
#### Signature: ``` func (q *Queue[T]) Size() int ```
#### Returns **int** value of the queue size.


### Tests


## Binary Search Tree
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



