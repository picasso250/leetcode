# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    # @param head, a ListNode
    # @return a boolean
    def hasCycle(self, head):
        saw = [head]
        while head is not None:
            if head.next is None:
                return False
            if head.next in saw:
                return True
        return False

print(Solution().hasCycle(None))

head = ListNode(1)
head.next = head
print(Solution().hasCycle(head))
