package tree

import (
    "fmt"
    "os"
    "io"
    "math"
    "dmitryro.com/queue"
)
// Data Structure to be used with BST nodes to handle tree traversal 
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
    Depth int
}

// Data Structure to return BST nodes with their depths
type NodeItem struct {
    Val int
    Depth int
}

// Data Structure to represent BST and its size
type BSTree struct {
    Size int
    Root *TreeNode
}

// A BST method to insert a new node in the tree
func (t *BSTree) InsertNode(Val int) *BSTree {
    if t.Root == nil {
        t.Root = &TreeNode{Val: Val, Left: nil, Right: nil}
    } else {
        t.Root.Insert(Val)
    }
    t.Size++
    return t
}

// A helper Insert method to be used with tree recursively traverse and append new value
func (n *TreeNode) Insert(Val int) {
    if n == nil {
        return
    } else if Val <= n.Val {
        if n.Left == nil {
            n.Left = &TreeNode{Val: Val, Left: nil, Right: nil}
        } else {
            n.Left.Insert(Val)
        }
    } else {
        if n.Right == nil {
            n.Right = &TreeNode{Val: Val, Left: nil, Right: nil}
        } else {
            n.Right.Insert(Val)
        }
    }
}

// Print the BST out starting at the root and attaching labels "L" for left, "R" for right sibling
func Print(w io.Writer, node *TreeNode, ns int, ch rune) {
    if node == nil {
        return
    }

    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
    Print(w, node.Left, ns+2, 'L')
    Print(w, node.Right, ns+2, 'R')
}

// Print the BST out using recurive helper
func (tree *BSTree) PrintTree() {
    Print(os.Stdout, tree.Root, 0, 'M')
}

// BST object method that uses helper to recursively travese tree and find node by value
func (tree *BSTree) FindNode(Val int) *NodeItem{
     node, found := tree.Root.Find(Val)
     var item = &NodeItem{Val, -1}

     if found {
         depth := NodeDepth(tree.Root, Val, 1)
         item = &NodeItem{node.Val, depth}
     } 
     return item
}
// A recursive helper used to traverse tree and find a node by value.
//  if found or not returns node itself and  true or false accordingly
func (t *TreeNode) Find(Value int) (TreeNode, bool) {

    if t == nil {
        return TreeNode{}, false
    }

    switch {
    case Value == t.Val:
        return *t, true
    case Value < t.Val:
        return t.Left.Find(Value)
    default:
        return t.Right.Find(Value)
    }
}

//DeleteNode Deletes the Item with Value from the tree
func (t *BSTree) DeleteNode(Value int) {
    t.Root.Delete(Value)
    t.Size--
}

// Helper used to recursively traverse and delete a node from tree
func (t *TreeNode) Delete(Value int) *TreeNode {

    if t == nil {
        return nil
    }

    if Value < t.Val {
        t.Left = t.Left.Delete(Value)
        return t
    }
    if Value > t.Val {
        t.Right = t.Right.Delete(Value)
        return t
    }

    if t.Left == nil && t.Right == nil {
        t = nil
        return nil
    }

    if t.Left == nil {
        t = t.Right
        return t
    }
    if t.Right == nil {
        t = t.Left
        return t
    }

    smallestValOnRight := t.Right
    for {
        //find smallest Value on the Right side
        if smallestValOnRight != nil && smallestValOnRight.Left != nil {
            smallestValOnRight = smallestValOnRight.Left
        } else {
            break
        }
    }

    t.Val = smallestValOnRight.Val
    t.Right = t.Right.Delete(t.Val)
    return t
}


// Search Tree - Checks whether the tree has a node
func (t *BSTree) TreeSearch(v int) bool {
    return t.Root.Search(v)
}


// Search - Checks whether the tree has a node
func (n *TreeNode) Search(v int) bool {
	if n == nil {
		return false
	}
	
	if n.Val < v {
		// move right
		return n.Right.Search(v)
	}else if n.Val > v {
		// move left
		return n.Left.Search(v)
	}
	return true
}

// Traverse BST to the downmost right node and find the maximal value, return when found
func (t *BSTree) FindTreeMax() int {
      current := t.Root
      var val = int(math.Inf(-1))

      if current != nil {
          for current != nil {
              current = current.Right           
              if current != nil {
                  val = current.Val          
              }
          }
      } 
      return val
}

//  Traverse BST to the downmost left node  and find the minimal value, return when found
func (t *BSTree) FindTreeMin() int {
      current := t.Root
      var val = int(math.Inf(1))

      if current != nil {
          for current != nil {
              current = current.Left
              if current != nil {
                  val = current.Val
              }
          }
      }
      return val
}

// Use helper BFS search and queue to find the list of deepest nodes in the tree
func (t *BSTree) DeepestNodes() []TreeNode {
    var bfsP =  findDeepestBFS(t.Root)
    bfs := *bfsP
    return bfs
}

// Construct a new tree from a list of integer values
func TreeFromArray(nodes []int) *BSTree{
    tree := &BSTree{}
    for i := 0; i < len(nodes); i++ {
       tree.InsertNode(nodes[i])
    }
    return tree
}

// Go over values in list and return false if found, true if found
func inNodes(nodes *[]TreeNode, v *TreeNode) bool {
    for i := 0; i < len(*nodes); i++ {
        if (*nodes)[i].Val == (*v).Val {
            return true
        }
    }
    return false
}

// Helper used to perform the actual discovery of deepest nodes
// Uses queue from queue package to do the BFS search for deepest nodes
// Returns slice with deepest nodes in it.
func findDeepestBFS(root *TreeNode) *[]TreeNode {
    q := new(queue.Queue[*TreeNode])
    q.Enqueue(root)
    q.Enqueue(root)
    var current = root
    var nodes []TreeNode  = make([]TreeNode, 0)   

    for q.Size() > 0 {
        v, valid := q.Dequeue()
        current = v
         
        if (q.Size() == 0) {
           break
        }
        for i := 0; i < q.Size(); i++ {
            if inNodes(&nodes, v) == false {
                nodes = append(nodes, (*v))
            }
        }

        if valid {
          
          if current.Left != nil {
             nodes = make([]TreeNode, 0)
             q.Enqueue(current.Left)            
          }

          if current.Right != nil{
             nodes = make([]TreeNode, 0)
             q.Enqueue(current.Right)
          }
          
        }        
    }    
    return &nodes
}

// Recursively traverse tree to find the depth of a node with value
// Return -1 if not found, positive integer number otherwise.
func NodeDepth(node *TreeNode, Val int, Depth int) int {
    if node == nil {
        return -1
    } else {
        lDepth := NodeDepth(node.Left, Val, Depth+1)
        rDepth := NodeDepth(node.Right, Val, Depth+1)
        if node.Val != Val {
            if lDepth > rDepth {
                return lDepth
            } else {
                return rDepth
            }
        } else {
            return Depth 
        }
    }
}

// Recursively traverse the tree and find the max depth of it
// Return -1 if empty, positive integer number for depth otherwise.
func MaxDepth(node *TreeNode) int {
    if node == nil {
        return -1
    } else {
        lDepth := MaxDepth(node.Left)
        rDepth := MaxDepth(node.Right)

        if lDepth > rDepth {
            return lDepth+1
        } else {
            return rDepth+1
        }
    }
}

// Return the current tree size
func (t *BSTree) TreeSize() int {
     return t.Size
}

// BST method to compute the max tree depth
// Uses recursive helper to traverse the tree from root
func (t *BSTree) MaxTreeDepth() int {
     depth := MaxDepth(t.Root)
     return depth
}

// BST method used to find the max depth and all the tree nodes
// Gets all the deepest nodes and attaches depth value to them
// Returns list of NodeItem objects that have values and depths.
func  (t *BSTree) DeepestNodeWithDepth() *[]NodeItem{
    var deepestItems []NodeItem = make([]NodeItem, 0)
    deepest := t.DeepestNodes()
    depth := MaxDepth(t.Root)
    for i:= 0; i < len(deepest); i++ {
       dp := NodeItem{Val:deepest[i].Val, Depth:depth}
       deepestItems = append(deepestItems, dp)
    }
    return &deepestItems
}

