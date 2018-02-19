package main

func sumNumberHelper(root *TreeNode,sum int) int {
    if root == nil {
        return 0
    } else if root.Left == nil && root.Right == nil {
        return sum*10 + root.Val
    } else {
        return sumNumberHelper(root.Left,sum*10 + root.Val) + 
               sumNumberHelper(root.Right,sum*10+root.Val)
    }
}
func sumNumbers(root *TreeNode) int {
    if root!=nil {
        return sumNumberHelper(root,0)
    }
    return 0
    
}

func main() {


}
