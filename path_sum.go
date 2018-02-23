package main

func pathSumHelper(root *TreeNode,intermediate []int,ans *[][]int,sum int) {
    
    if root!=nil {
        
        sum-=root.Val
        intermediate=append(intermediate,root.Val)
        pathSumHelper(root.Left,intermediate,ans,sum)
        pathSumHelper(root.Right,intermediate,ans,sum)
        
        if root.Left == nil && root.Right == nil && sum==0 {
            *ans=append(*ans,intermediate)
            fmt.Println(*ans)
            
        }
    }
}
func pathSum(root *TreeNode, sum int) [][]int {
    
    intermediate:=make([]int,0)
    ans:=make([][]int,0)
    pathSumHelper(root,intermediate,&ans,sum)
    return ans
    
}

func main() {



}
