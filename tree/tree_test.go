package tree

import (
    "fmt"
    "testing"
)

// Test inserting a new value to BST
// Create a new BST from the list of 6 values
// Insert a new node with value 15 to the tree
// Assert success if new tree size is 7, assert failure othewise
func TestInsert(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   tree.InsertNode(15)
   
   size := tree.TreeSize()
   if size == 7 {
       fmt.Printf("Successfully inserted new node - new size is  %d\n", size)
   } else {
       t.Errorf("Failed to insert new node - the size is %d\n", size)
   }   

   fmt.Println("Insert")
}

// Test deleting a value from BST
// Create a new BST from the list of 6 values
// Delete a node with value 90 from the tree
// Assert success if new tree size is 5, assert failure othewise
func TestDelete(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   tree.DeleteNode(90)

   size := tree.TreeSize()
   if size == 5 {
       fmt.Printf("Successfully deleted a node - new size is  %d\n", size)
   } else {
       t.Errorf("Failed to delete a node - the size is %d\n", size)
   }

   fmt.Println("Delete")
}

// Test valid sizing of BST
// Create a new BST from the list of 6 values and compute size
// Assert success if tree size is 6, assert failure othewise
func TestTreeSize(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   size := tree.TreeSize()

   if size == 6 {
       fmt.Printf("Successfully calculated tree size - %d\n", size)
   } else {
       t.Errorf("Failed to calculate tree size - %d\n", size)
   }
}

// Test returning all the nodes from the deepest level
// Create a new BST with 6 nodes in it
// Read the deepest nodes of the tree
// If the call returns 2 elements - assert success, if no - assert failure
func TestDeepest(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   deepestNodes := tree.DeepestNodeWithDepth()
   lenDeepest := len(*deepestNodes)
   if lenDeepest == 2 {
       fmt.Printf("Successfully found deepest nodes - %d\n", lenDeepest)
   } else {
       t.Errorf("Failed to find deepest nodes - %d\n", lenDeepest)
   }
   fmt.Println("Deepest")
}


// Test computing the maximal depth of the tree
// Create a test tree from list of values specified
// Compute the depth - assert success if the depth is 3, assert failure otherwise
func TestMaxDepth(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   depth := tree.MaxTreeDepth()
   if depth == 3 {
       fmt.Printf("Successfully computed the max tree depth - %d\n", depth)
   } else {
       t.Errorf("Failed to compute the max tree depth %d - got %d\n", 3, depth)
   }

   fmt.Println("Max Depth")
}

// Test finding a node with depth in BST
// Create a test tree from list of values specified
// Find node with value 82 and assert success if value found and depth is 3, assert failure otherwise
func TestFindNode(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   node := tree.FindNode(82)

   if node.Val == 82 && node.Depth == 3 {
       fmt.Printf("Successfully found node and its depth in the tree - %d - %d\n", node.Val, node.Depth)
   } else {
       t.Errorf("Failed to find node and its depth in the tree - %d - %d\n", node.Val, node.Depth) 
   }

   fmt.Println("Find")
}

// Test searching for an element in BST
// Create a test tree from list of values specified
// Value 82 is in the tree, value 32 is not
// For value 82 assert success if found, failure if not found
// For value 32 assert success if not found, failure if found
func TestTreeSearch(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   found := tree.TreeSearch(82)

   if found == true {
       fmt.Printf("Successfully searched tree for element %d - it was found.\n", 82)
   } else {
       t.Errorf("Failed searching the tree for element %d - it was not found\n", 82)
   }
   
   more_found := tree.TreeSearch(32)
   if more_found == false {
       fmt.Printf("Successfully searched tree for element %d - it was never inserted and was not found.\n", 32)
   } else {
       t.Errorf("Failed searching the tree for element %d - it was never inserted but was found\n", 32)
   }
   

   fmt.Println("Find")
}

// Test finding minimal value in BST
// Construct a test tree from the list of values
// Assert success if min value is 7, assert failure otherwise
func TestFindTreeMin(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)

   if tree.FindTreeMin() == 7 {
       fmt.Printf("Successfully found minimum of the tree - %d\n", 7)
   } else {
       t.Errorf("Failed to find the minimum of the tree - %d\n", 7)
   }
   fmt.Println("Find Min")
}

// Test finding maximal value in BST
// Construct a test tree from the list of values
// Assert success if max value is 90, assert failure otherwise
func TestFindTreeMax(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)

   if tree.FindTreeMax() == 90 {
       fmt.Printf("Successfully found maximum of the tree - %d\n", 90)
   } else {
       t.Errorf("Failed to find the maximum of the tree - %d\n", 90)
   }
   fmt.Println("Find Max")
}

// Test constructing BST from a list of values
// Construct two test trees
// For the first tree, assert success if depth is 3, assert failure otherwise
// For the second tree, assert success if depth is 2, assert failure othewise
func TestTreeFromArray(t *testing.T) {
   l1 := []int{12, 11, 90, 82, 7, 9}
   t1 := TreeFromArray(l1)

   if t1.MaxTreeDepth() == 3 {
       fmt.Printf("Successfully created a tree with depth - %d\n", 3)
   } else {
       t.Errorf("Failed creating a tree with depth - %d\n", 3)
   }

   l2 := []int{26, 82, 16, 92, 33}
   t2 := TreeFromArray(l2)

   if t2.MaxTreeDepth() == 2 {
       fmt.Printf("Successfully created a tree with depth - %d\n", 2)
   } else {
       t.Errorf("Failed creating a tree with depth - %d\n", 2)
   }

   fmt.Println("Tree From Array")
}
