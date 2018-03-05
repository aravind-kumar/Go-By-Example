package main
type TreeNode struct {
    Val int
    Left,Right *TreeNode
}

func findTargetHelper(root *TreeNode,k int,set *map[int]bool) bool {
    if root!=nil {
        if _,ok := ((*set)[k-root.Val]) ; ok {
            return true
        } 
        (*set)[root.Val] = true
        return findTargetHelper(root.Left,k,set) || findTargetHelper(root.Right,k,set)
    }   
    return false
}

func findTarget(root *TreeNode, k int) bool {
    if root!=nil {
        hashset := make(map[int]bool,0)
        return findTargetHelper(root,k,&hashset)
    }
    return k==0;
    
}

func main() {



}
