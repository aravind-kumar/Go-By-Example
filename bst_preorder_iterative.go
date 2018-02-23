package main

func preorderTraversal(root *TreeNode) []int {
    
    if root != nil {
        stack := make([](*TreeNode),0)
        stack = append(stack,root)
        output := make([]int,0)
        
        for len(stack) > 0 {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            
            output=append(output,(top.Val))
            
            if top.Right != nil {
                stack=append(stack,top.Right)
            }
            if top.Left != nil {
                stack=append(stack,top.Left)
            }
        }
        
        return output
    }
    return make([]int,0)
    
    
}


func main() {



}
