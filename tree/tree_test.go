package tree

import (
    "fmt"
    "testing"
)

func TestInsert(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   tree.Insert(15)
   
   size := tree.TreeSize()
   if size == 7 {
       fmt.Printf("Successfully inserted new node - new size is  %d\n", size)
   } else {
       fmt.Printf("Failed to insert new node - the size is %d\n", size)
   }   

   fmt.Println("Insert")
}

func TestDelete(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   tree.Delete(90)

   size := tree.TreeSize()
   if size == 5 {
       fmt.Printf("Successfully deleted a node - new size is  %d\n", size)
   } else {
       fmt.Printf("Failed to delete a node - the size is %d\n", size)
   }

   fmt.Println("Delete")
}

func TestTreeSize(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   size := tree.TreeSize()

   if size == 6 {
       fmt.Printf("Successfully calculated tree size - %d\n", size)
   } else {
       fmt.Printf("Failed to calculate tree size - %d\n", size)
   }
}

func TestDeepest(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   deepestNodes := tree.DeepestNodeWithDepth()
   lenDeepest := len(*deepestNodes)
   if lenDeepest == 2 {
          fmt.Printf("Successfully found deepest nodes - %d\n", lenDeepest)
   } else {
          fmt.Printf("Failed to find deepest nodes - %d\n", lenDeepest)
   }
   fmt.Println("Deepest")
}

func TestMaxDepth(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   depth := tree.MaxTreeDepth()
   if depth == 3 {
        fmt.Printf("Successfully computed the max tree depth - %d\n", depth)
   } else {
        fmt.Printf("Failed to compute the max tree depth %d - got %d\n", 3, depth)
   }

   fmt.Println("Max Depth")
}

func TestFindNode(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)
   node := tree.FindNode(82)

   if node.Val == 82 && node.Depth == 3 {
      fmt.Printf("Successfully found node and its depth in the tree - %d - %d\n", node.Val, node.Depth)
   } else {
      fmt.Printf("Failed to find node and its depth in the tree - %d - %d\n", node.Val, node.Depth) 
   }

   fmt.Println("Find")
}

func TestFindTreeMin(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)

   if tree.FindTreeMin() == 7 {
         fmt.Printf("Successfully found minimum of the tree - %d\n", 7)
   } else {
         fmt.Printf("Failed to find the minimum of the tree - %d\n", 7)
   }
   fmt.Println("Find Min")
}

func TestFindTreeMax(t *testing.T) {
   l := []int{12, 11, 90, 82, 7, 9}
   tree:= TreeFromArray(l)

   if tree.FindTreeMax() == 90 {
         fmt.Printf("Successfully found maximum of the tree - %d\n", 90)
   } else {
         fmt.Printf("Failed to find the maximum of the tree - %d\n", 90)
   }
   fmt.Println("Find Max")
}

func TestTreeFromArray(t *testing.T) {
   l1 := []int{12, 11, 90, 82, 7, 9}
   t1 := TreeFromArray(l1)

   if t1.TreeDepth() == 3 {
        fmt.Printf("Successfully created a tree with depth - %d\n", 3)
   } else {
        fmt.Printf("Failed creating a tree with depth - %d\n", 3)
   }

   l2 := []int{26, 82, 16, 92, 33}
   t2 := TreeFromArray(l2)

   if t2.TreeDepth() == 2 {
       fmt.Printf("Successfully created a tree with depth - %d\n", 2)
   } else {
       fmt.Printf("Failed creating a tree with depth - %d\n", 2)
   }

   fmt.Println("Tree From Array")
}
