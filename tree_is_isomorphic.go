package main


type TreeNode struct {
    val int
    left,right *TreeNode
}

func isIsomorphicHelper(left,right *TreeNode) bool {

    if left == nil && right == nil {
       return true 
    } else if left == nil || right == nil || left.Val != right.Val {
       return false 
    }
    
    return (isIsomorphicHelper(left.left,right.left) && isIsomorphicHelper(left.right,right.right))
           ||
           (isIsomorphicHelper(left.right,right.left) && isIsomorphicHelper(left.left,right.right))
}

func main() {


}

