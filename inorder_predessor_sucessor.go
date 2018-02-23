package main
import "fmt"


type TreeNode struct {
    val int
    left,right  *TreeNode
}

func getInorderPredessor(root,pre *TreeNode) *TreeNode {
    if root!=nil && pre!=nil {
       if pre.val <= root.val {
           return getInorderPredessor(root.left,pre) 
       } else {
          pred := getInorderPredessor(root.right,pre)
          if pred == nil {
              return root 
          } else {
             return pred 
          }
       }
    }
    return nil
}

func getInorderSucessor(root,pre *TreeNode) *TreeNode {
    if root!=nil && pre!=nil {
        if pre.val >= root.val {
            return getInorderSucessor(root.right,pre) 
        } else {
            sucessor := getInorderSucessor(root.left,pre)
            if sucessor == nil {
                return root 
            } else {
                return nil 
            }
        }
    }
    return nil
}

func main() {


}
