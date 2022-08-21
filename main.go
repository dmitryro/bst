package main

import (
      "fmt"
      "dmitryro.com/tree"
)

func printTree(l []int) {
    t := tree.TreeFromArray(l)
    t.PrintTree()
    dn := t.DeepestNodeWithDepth()
    for i := 0; i < len(*dn); i++ {
        fmt.Printf("\nDepth %d Value %d\n", (*dn)[i].Depth, (*dn)[i].Val)
    }
}

func main() {
    l1 := []int{12, 11, 90, 82, 7, 9} 
    printTree(l1)


    l2 := []int{26, 82, 16, 92, 33}
    printTree(l2)

    t := tree.TreeFromArray(l1)
    node := t.FindNode(82)

    fmt.Printf("NODE WAS %d DEPTH WAS %d", node.Val, node.Depth)
}
