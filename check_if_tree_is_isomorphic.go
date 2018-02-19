package main
import "fmt"

Type Node struct {
    Val int
    left,right *Node
}


func isSymmetricHelper(left,right *Node) bool {
    if left == nil && right == nil {
        return true 
    }
    if left!=nil && right!=nil && left.Val == right.Val {
       return isSymmetricHelper(left.left,right.right) && isSymmetricHelper(left.right,right.left)
    }
    return false
}

func isSymmetricTree(root *Node) bool {
    if root != nil {
        return isSymmetricHelper(root.Left,root.Right)    
    }
    return false
}

func main() {

}
