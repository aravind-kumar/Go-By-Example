package main

type TreeNode struct {
    data int
    left,right *TreeNode
}


func LCA(root,p,q *TreeNode) *TreeNode {

     if root == nil {
        return nil  
     } else if root == p || root == q {
        return root 
     } 
     left,right := LCA(root.left,p,q),LCA(root.right,p,q)

     if left!=nil && right!=nil {
        return root
     } else if left == nil {
        return right 
     } else {
        return left 
     }
}


func main() {



}
