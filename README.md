# Go Implementation of Binary Search Tree
Binary Search Tree Implementation using Go, Queue and Tree data structures, Breadth First Search for tree traversal.

## Getting Started
To run the project, clone it to your loccal environment:
```
  git clone git@github.com:dmitryro/bst.git
```
Please have Go 1.18 or later installed, if not installed - please follow the instructions in the next section.
Please refer to this README document or comments within code if additional clarifications are needed.


## Installing Go 1.18 or later
Please follow the [instructions](https://www.digitalocean.com/community/tutorial_collections/how-to-install-go) based on
what system you use.


## Initial Assumptions
To successfully run the project some initial assumptions need to be made:

**Go 1.18 SDK** or later must be installed in your system.

**GOROOT** environment vairable should point to the proper Go SDK location.

**GOPATH** environment variable should be pointing to  a valid workspace.

**Queue** data structure is used to implement **BFS** search. We use a custom implementation, while alternative 
      third party implementations exist.

**Tree** and **Queue** behaviors are covered by unit tests, to run which please follow the instructions below.

**Capitalized Method and Function Naming Notation** will be used to expose behaviors outside package.
     

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
#### [**DeepestNodeWithDepth**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L329):
  - **Description**: **BST** Object method finding the largest depth and finding the nodes at this deepest level.

  - **Signature**: ``` func  (t *BSTree) DeepestNodeWithDepth() *[]NodeItem  ```

  - **Returned Value**: Returns a list of **NodeItem** objects with values and their depths found at the deepest level of the tree.


#### [**DeepestNodes**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L214):
  - **Description**: **BST** Object method finding the nodes with the larges depth using DFS helper call utilizing queue structure.

  - **Signature**: ``` func (t *BSTree) DeepestNodes() []TreeNode ```

  - **Returned Value**: Returns a list of **TreeNode** objects found at the deepest level of the tree that were fetched by **findDeepestBFS** helper.


#### [**DeleteNode**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L108):
  - **Description**: **BST** Object method allowing to add a new node to **BST**.

  - **Signature**: ``` func (t *BSTree) DeleteNode(Value int)  ```

  - **Returned Value**: Returns no values


#### [**Delete**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L114):
  - **Description**: Node Object method receives the tree root and does the actual recurive traversal of the BST. It's used as helper method by tree to insert a new node to *BST**.

  - **Signature**: ``` func (t *TreeNode) Delete(Value int) *TreeNode  ``` - the tree root is provided as initial parameter to traverse.

  - **Returned Value**: Returns no values


#### [**findDeepestBFS**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L240):
  - **Description**: **BST** Object method allowing **BFS** search for deepest nodes using **Queue** structure.

  - **Signature**: ``` func findDeepestBFS(root *TreeNode) *[]TreeNode  ```

  - **Returned Value**: Returns **[]TreeNode** list of pointers to the deepest nodes.


#### [**FindNode**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L79):
  - **Description**: **BST** Object method allowing to add a new node to **BST**.

  - **Signature**: ``` func (tree *BSTree) FindNode(Val int) *NodeItem  ```

  - **Returned Value**: Returns no values


#### [**Find**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L90):
  - **Description**: Node Object method receives the tree root and does the actual recurive traversal of the BST. It's used as helper method by tree to insert a new node to *BST**.

  - **Signature**: ``` func (t *TreeNode) Find(Value int) (TreeNode, bool)  ``` - the tree root is provided as initial parameter to traverse.

  - **Returned Value**: Returns the node if found, empty record if not, true/false boolean depending succes or falure of the search.


#### [**FindTreeMax**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L182):
  - **Description**: Use the existing tree's root to go to the rightmost downside element and return the value. 

  - **Signature**: ``` func (t *BSTree) FindTreeMax() int ``` - The tree root from tree object is used to traverse, recursively traversed by helper.

  - **Returned Value**: Returns the maximal value if found


#### [**FindTreeMin**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L198):
  - **Description**: Use the existing tree's root to go to the leftmost downside element and return the value.

  - **Signature**: ``` func (t *BSTree) FindTreeMin() int   ``` - The tree root from tree object is used to traverse, recursively traversed by helper.

  - **Returned Value**: Returns the minimal value if found


#### [**InsertNode**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L30):
  - **Description**: **BST** Object method allowing to add a new node to **BST**.

  - **Signature**: ``` func (t *BSTree) InsertNode(Val int) *BSTree  ```

  - **Returned Value**: Returns no values


#### [**Insert**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L41):
  - **Description**: Node Object method receives the tree root and does the actual recurive traversal of the BST. It's used as helper method by tree to insert a new node to *BST**.

  - **Signature**: ``` func (n *TreeNode) Insert(Val int)  ``` - the tree root is provided as initial parameter to traverse.

  - **Returned Value**: Returns no values


#### [**MaxTreeDepth**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L322):
  - **Description**: **BST** Object method allowing to get the starting node (root), value to check and return depth at which the value is located.

  - **Signature**: ``` func (t *BSTree) MaxTreeDepth() int  ```

  - **Returned Value**: Depth at which the value was found in **BST**.


#### [**MaxDepth**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L300):
  - **Description**: **BST** Object method allowing to get the starting node (root), value to check and return depth at which the value is located.

  - **Signature**: ``` func MaxDepth(node *TreeNode) int ```

  - **Returned Value**: Depth at which the value was found in **BST**.


#### [**NodeDepth**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L280):
  - **Description**: **BST** Object method allowing to get the starting node (root), value to check and return depth at which the value is located.

  - **Signature**: ``` func NodeDepth(node *TreeNode, Val int, Depth int) int ```

  - **Returned Value**: Depth at which the value was found in **BST**.


#### [**inNodes**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L230):
  - **Description**: Helper function to check if the value is found in slice or not.

  - **Signature**: ``` func inNodes(nodes *[]TreeNode, v *TreeNode) bool  ```

  - **Returned Value**: Returns boolean - **true** if value was found in slice, **false** otherwise.


#### [**PrintTree**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L74):
  - **Description**: **BST** Object method allowing to print the entire tree by using the helper call to recursively traverse the tree.

  - **Signature**: ``` func (tree *BSTree) PrintTree() ```

  - **Returned Value**: Returns no values


#### [**Print**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L60):
  - **Description**: Print the tree by recurively traversing it starting at the root level.

  - **Signature**: ``` func Print(w io.Writer, node *TreeNode, ns int, ch rune)  ``` - gets the tree root, start at level 0 and mark it as 'M' - middle.

  - **Returned Value**: Returns no values


#### [**TreeSearch**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L160):
  - **Description**: **BST** Object method allowing search for a value in the tree and return true or false depending on outcome.

  - **Signature**: ``` func (t *BSTree) TreeSearch(v int) bool ``` - receives the value to find.

  - **Returned Value**: Returns **true** if value was found, **false** otherwise.


#### [**Search**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L166):
  - **Description**: Helper functionality to recursively search tree and return true when value is found, false otherwise.

  - **Signature**: ``` func (n *TreeNode) Search(v int) bool   ``` - receives the value to find.

  - **Returned Value**: Returns **true** if value was found, **false** otherwise.


#### [**TreeFromArray**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L222):
  - **Description**: **BST** Object method allowing search for a value in the tree and return true or false depending on outcome.

  - **Signature**: ``` func TreeFromArray(nodes []int) *BSTree ``` - receives the slice of values to construct tree.

  - **Returned Value**: Returns the pointer to the root of the newly created **BST**.


#### [**TreeSize**](https://github.com/dmitryro/gotests/blob/master/tree/tree.go#L318):
  - **Description**: **BST** Object method allowing search for a value in the tree and return true or false depending on outcome.

  - **Signature**: ``` func (t *BSTree) TreeSize() int  ``` - receives the value to find.

  - **Returned Value**: Returns the current size of the **BST**.

## Running the project
### Running the main program
Once cloned and verified the Go version 1.18 or later is installed, run
```
go run main.go
```
This program will create two sample trees out of lists of values, construct them and search for deepest nodes, 
print them out with their correpsonding depths.

### Functions
#### [**printTree**](https://github.com/dmitryro/gotests/blob/master/main.go#L8):
  - **Description**: The function initiates new Tree object, prints the generated tree out and prints the deepest nodes of the generated tree.

  - **Signature**: ``` func printTree(l []int)  ``` - A list of values is provided

  - **Returned Value**: Returns no values



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



