package main


type Node struct {
    val int
    Next *Node
}

func reverseLL(head *Node) *Node {
     
     var pre *Node = nil
     
     for head != nil {
         var newHead = head.Next
         head.Next = pre
         pre = head
         head = newHead
     }
     return pre
}


func main() {


}
