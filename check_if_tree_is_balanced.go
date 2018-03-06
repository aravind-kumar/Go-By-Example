package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func abs(a int) int {
    if a<0 {
        return -a
    } else {
        return a
    }
}

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func height(root *TreeNode) int {
    if root != nil {
        lHeight := height(root.Left)
        if lHeight == -1 {
            return -1
        }
        rHeight := height(root.Right)
        if rHeight == -1 {
            return -1
        }
        if abs(lHeight-rHeight) > 1 {
            return -1
        }
        return max(lHeight,rHeight)+1
    }
    return 0
}

func isBalanced(root *TreeNode) bool {
    return height(root) != -1
}

func main() {




}
