package main

type Node struct {
    val int
    left,right *Node 
}

func LCA(root *Node,node1,node2 Node) *Node {
   for root!=nil {
        if root.val < node1.val && root.val < node2.val {
            root = root.right 
        } else if root.val > node1.val && root.val > node2.val {
            root = root.left 
        } else {
            break 
        }
    }
    return root

}

func main() {



}
