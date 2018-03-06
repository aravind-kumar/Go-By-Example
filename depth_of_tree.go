package main

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func maxDepth(root *TreeNode) int {
    if root != nil {
        lDepth := maxDepth(root.Left)
        rDepth := maxDepth(root.Right)
        return max(lDepth,rDepth)+1
    }
    return 0
    
}

func main() {



}
