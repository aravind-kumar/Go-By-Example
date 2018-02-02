func hasCycle(ListNode *head) {
        ListNode *current=head,*runner=head;
        while(current && runner && current->next != nullptr) {
            current = current->next->next;
            runner = runner->next;
            if (current == runner)
                return true;
        }
        return false;
    }
