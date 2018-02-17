package main

func isValidBSTHelper(root,pre *TreeNode) bool {
    if root!=nil {
        if !isValidBSTHelper(root.Left,pre) {
            return false
        }
        if pre!=nil && pre.Val >= root.Val {
            return false
        }
        pre = root
        if !isValidBSTHelper(root.Right,pre) {
            return false
        }
    }   
    return true
}

func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root,nil)
}

func main() {


}
