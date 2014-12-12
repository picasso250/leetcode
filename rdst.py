# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
    def __repr__(self):
        r = str(self.val)+'->'
        if self.next:
            r += repr(self.next)
        return r

class Solution:
    # @param head, a ListNode
    # @return a ListNode
    def deleteDuplicates(self, head):
        if head is None:
            return head
        parent = head
        cur = head.next
        while cur is not None:
            assert parent.next == cur
            if parent.val == cur.val:
                cur = cur.next
                # parent stay
                parent.next = cur
            else:
                parent = cur
                cur = cur.next
            # assert cur.val != parent.val
        return head
            
# test
head = ListNode(1)
head.next = ListNode(1)
head.next.next = ListNode(1)

print(head)
print(Solution().deleteDuplicates(head))
