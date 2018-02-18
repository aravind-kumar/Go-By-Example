package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a,b int) int {
    if a>b {
        return a
    } else {
        return b
    }
}

func height(root *TreeNode,diameter *int) int {
    if root == nil {
        return 0
    } else {
        var lheight int = height(root.Left,diameter)
        var rheight int = height(root.Right,diameter)
        
        *diameter = max(*diameter,lheight+rheight)
        
        return max(lheight,rheight)+1
    }
}

func diameterOfBinaryTree(root *TreeNode) int {
    
    if root != nil {
        diameter := math.MinInt32        
        height(root,&diameter)
        return diameter
    }
    return 0
    
}


func main() {



}
