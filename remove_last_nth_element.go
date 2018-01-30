func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    var current,runner *ListNode = head,head
    
    var temp *ListNode
    var i int
    for i=0 ; i<n ; i+=1 {
        current=current.Next
    }
    
    if current!=nil {
        for current.Next != nil {
            runner = runner.Next
            current = current.Next
        }
        temp = runner.Next
        runner.Next = temp.Next
    } else {
        temp = head
        head = head.Next
    } 
    temp.Next = nil
    temp = nil
    return head
}

