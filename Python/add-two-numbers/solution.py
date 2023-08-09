# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        res_start = l1
        zeroNode = ListNode(0)
        while True:
            number = l1.val + l2.val
            l1.val = number % 10
            if l1.next == None and l2.next == None:
                if number >= 10:
                    l1.next = ListNode(val=1)
                return res_start
            if l1.next and l2.next:
                l1 = l1.next
                l2 = l2.next
            elif l1.next:
                l1 = l1.next
                l2 = zeroNode
            else:
                l1.next = l2.next
                l1 = l1.next
                l2 = zeroNode
            if number >= 10: l1.val += 1
