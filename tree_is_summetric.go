func isSymmetricHelper(left,right *TreeNode) bool {
    if left==nil && right==nil {
        return true
    } else if left!=nil && right!=nil && left.Val == right.Val {
        return isSymmetricHelper(left.Left,right.Right) && isSymmetricHelper(left.Right,right.Left)
    }
    return false
}

func isSymmetric(root *TreeNode) bool {
    if root != nil {
        return isSymmetricHelper(root.Left,root.Right)
    }
    return true
}
