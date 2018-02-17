package main

func reverseList(head *ListNode) *ListNode {
    if head==nil || head.Next == nil {
        return head
    } else {
        temp := reverseList(head.Next)
        head.Next.Next = head
        head.Next = nil
        return temp
    }
}

func main() {

}
