#
# @lc app=leetcode id=24 lang=python3
#
# [24] Swap Nodes in Pairs
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def swapPairs(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        GuardNode = p = ListNode(0)
        GuardNode.next = head
        
        while head and head.next:
            tem = head.next        
            head.next = tem.next 
            tem.next = head   
            p.next = tem     
            p = head               
            head = head.next 
        
        return GuardNode.next
            
# @lc code=end

