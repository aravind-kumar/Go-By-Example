func deleteNode(node *ListNode) {
        if node!=nil {
            var nextNode *ListNode = node.next;
            nextNode.val,node.val = node.val,nextNode.val
            node.next = nextNode.next;
            nextNode.next = nullptr;
            nextNode = nil
        }
}
