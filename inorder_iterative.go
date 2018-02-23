package main

func addAllLeft(root *TreeNode,stack *[](*TreeNode)) {
    for root != nil {
        *stack=append(*stack,root)
        root = root.Left
    }    
}

func inorderTraversal(root *TreeNode) []int {
    
    if root!=nil {
        stack:=make([](*TreeNode),0)
        
        output := make([]int,0)
        
        addAllLeft(root,&stack)
        
        for len(stack) > 0 {
            
            top:=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            
            output = append(output,top.Val)
            
            addAllLeft(top.Right,&stack)
    
        }
        return output
    }
    return make([]int,0)
    
}

func main() {



}
