# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        if not list1 and not list2: return list1
        elif not list2: return list1
        elif not list1: return list2

        current_node, other_node = [list1, list2] if list1.val <= list2.val else [list2, list1]
        start_node = current_node
        while current_node.next:
            if not other_node:
                return start_node
            elif current_node.val <= other_node.val < current_node.next.val:
                next_to_visit = other_node.next
                other_node.next = current_node.next
                current_node.next = other_node
                other_node = next_to_visit

            if current_node.next:
                current_node = current_node.next
            else:
                return start_node

        if other_node:
            current_node.next = other_node
            return start_node
        else:
            current_node.next = other_node
            return start_node       
