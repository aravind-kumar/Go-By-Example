func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    var newHead,tail *ListNode;
    for l1 != nil || l2 != nil {
        var smallerNode *ListNode;
        
        if l1 != nil && l2 != nil {
            if l1.Val < l2.Val {
                smallerNode = l1
                l1 = l1.Next
            } else {
                smallerNode = l2
                l2 = l2.Next
            }
        } else if l1 !=nil {
            smallerNode = l1
            l1 = l1.Next
        } else {
            smallerNode = l2
            l2 = l2.Next
        }
        
        if newHead == nil {
            newHead = smallerNode
            tail = newHead
        } else {
            tail.Next = smallerNode
            tail = smallerNode
        }
    }
    return newHead;
}
