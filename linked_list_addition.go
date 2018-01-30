func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var carry int = 0
    var head,tail *ListNode
    for l1!=nil || l2!=nil || carry>0 {
        var x,y int = 0,0
        
        if l1!=nil {
            x = l1.Val
        }
        
        if l2!=nil {
            y = l2.Val
        }
        
        var sum int = x+y+carry
        
        var finalSum = sum % 10
        carry = sum/10
        
        if head == nil {
            head = new (ListNode)
            tail = head
        } else {
            tail.Next = new (ListNode)
            tail=tail.Next;
        }
        tail.Val = finalSum
        
        if l1 != nil {
            l1 = l1.Next    
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    
    return head
