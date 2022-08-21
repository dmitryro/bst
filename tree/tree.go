package tree

import (
    "fmt"
    "os"
    "io"
    "math"
    "dmitryro.com/queue"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
    Depth int
}


type NodeItem struct {
    Val int
    Depth int
}


type BSTree struct {
    Size int
    Root *TreeNode
}

func (t *BSTree) Insert(Val int) *BSTree {
    if t.Root == nil {
        t.Root = &TreeNode{Val: Val, Left: nil, Right: nil}
    } else {
        t.Root.Insert(Val)
    }
    t.Size++
    return t
}

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

func print(w io.Writer, node *TreeNode, ns int, ch rune) {
    if node == nil {
        return
    }

    for i := 0; i < ns; i++ {
        fmt.Fprint(w, " ")
    }
    fmt.Fprintf(w, "%c:%v\n", ch, node.Val)
    print(w, node.Left, ns+2, 'L')
    print(w, node.Right, ns+2, 'R')
}

func (tree *BSTree) Print() {
    print(os.Stdout, tree.Root, 0, 'M')
}

func (tree *BSTree) FindNode(Val int) *NodeItem{
     node, found := tree.Root.Find(Val)
     var item = &NodeItem{Val, -1}

     if found {
         depth := nodeDepth(tree.Root, Val, 1)
         item = &NodeItem{node.Val, depth}
     } 
     return item
}

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

//Delete removes the Item with Value from the tree
func (t *BSTree) Delete(Value int) {
    t.Root.remove(Value)
    t.Size--
}

func (t *TreeNode) remove(Value int) *TreeNode {

    if t == nil {
        return nil
    }

    if Value < t.Val {
        t.Left = t.Left.remove(Value)
        return t
    }
    if Value > t.Val {
        t.Right = t.Right.remove(Value)
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
    t.Right = t.Right.remove(t.Val)
    return t
}


// Search - Checks whether the tree has a node
func (n *TreeNode) Search(v int)bool{
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


//FindMax finds the max element in the given BST
func (root *TreeNode) FindMax() int {
    if root == nil {
        minusInf := int(math.Inf(-1))
        root = &TreeNode{Val: minusInf, Left: nil, Right: nil}     
    }
    val := root.Val

    lLeft := root.Left.FindMax() 
    lRight := root.Right.FindMax()

    if lLeft > val {
        val = lLeft
    }

    if lRight > val {
        val = lRight
    }
    return val
}


func (t *BSTree) TreeDepth() int {
    return maxDepth(t.Root)
}

func (t *BSTree) DeepestNodes() []TreeNode {
    var bfsP =  findDeepestBFS(t.Root)
    bfs := *bfsP
    return bfs
}

func TreeFromArray(nodes []int) *BSTree{
    tree := &BSTree{}
    for i := 0; i < len(nodes); i++ {
       tree.Insert(nodes[i])
    }
    return tree
}

func notInNodes(nodes *[]TreeNode, v *TreeNode) bool {
    for i := 0; i < len(*nodes); i++ {
        if (*nodes)[i].Val == (*v).Val {
            return false
        }
    }
    return true
}

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
            if notInNodes(&nodes, v) == true {
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

func nodeDepth(node *TreeNode, Val int, Depth int) int {
    if node == nil {
        return -1
    } else {
        lDepth := nodeDepth(node.Left, Val, Depth+1)
        rDepth := nodeDepth(node.Right, Val, Depth+1)
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

func maxDepth(node *TreeNode) int {
    if node == nil {
        return -1
    } else {
        lDepth := maxDepth(node.Left)
        rDepth := maxDepth(node.Right)

        if lDepth > rDepth {
            return lDepth+1
        } else {
            return rDepth+1
        }
    }
}

func (t *BSTree) TreeSize() int {
     return t.Size
}

func (t *BSTree) MaxTreeDepth() int {
     depth := maxDepth(t.Root)
     return depth
}

func  (t *BSTree) DeepestNodeWithDepth() *[]NodeItem{
    var deepestItems []NodeItem = make([]NodeItem, 0)
    deepest := t.DeepestNodes()
    depth := maxDepth(t.Root)
    for i:= 0; i < len(deepest); i++ {
       dp := NodeItem{Val:deepest[i].Val, Depth:depth}
       deepestItems = append(deepestItems, dp)
    }
    return &deepestItems
}

