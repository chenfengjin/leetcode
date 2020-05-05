#
# @lc app=leetcode id=160 lang=python3
#
# [160] Intersection of Two Linked Lists
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        if not headA or not headB:
            return None
        currentA = headA
        currentB = headB
        while headA.next:
            headA = headB.next
        while headB.next:
            headB = headB.next
        if not headB== headA:
            return False
        while not currentA == currentB:
            if currentA.next:
                currentA = currentA.next
            else:
                currentA = headB
            if currentB.next:
                currentB = currentB.next
            else:
                currentB = headA
        return currentB

        
# @lc code=end

