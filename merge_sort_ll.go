func merge(left *ListNode, right *ListNode) *ListNode {
    if left == nil && right == nil {
        return nil
    }else {
        var head,tail *ListNode
        for left != nil && right != nil {
            var temp *ListNode
            if left.Val < right.Val {
                temp = left
                left = left.Next
            }else {
                temp = right
                right = right.Next
            }
            if head == nil {
                head,tail = temp,temp
            } else {
                tail.Next = temp
                tail = tail.Next
            }
        }
        if left != nil {
            tail.Next = left
        } else {
            tail.Next = right
        }
        return head
    }
}

func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var current,runner *ListNode = head.Next,head
    
    for current != nil && current.Next != nil {
        runner = runner.Next
        current = current.Next.Next
    }
    
    current = runner.Next
    runner.Next = nil
    
    return merge(sortList(head),sortList(current))
}
