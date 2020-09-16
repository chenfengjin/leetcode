#
# @lc app=leetcode id=143 lang=python3
#
# [143] Reorder List
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
#  一个更快的算法
#  只入栈一半，然后一次遍历，这样的话可以省一半空间，时间比较难算
class Solution:
    def reorderList(self, head: ListNode) -> None:
        """
        Do not return anything, modify head in-place instead.
        """
        if not head:
            return 
        s = []
        current = head.next
        while current:
            s.append(current)
            current = current.next   
        current = head
        flag = True
        while s:
            if flag:
                current.next = s.pop()
            else:
                current.next = s.pop(0)
            current = current.next
            flag  = not flag
        current.next = None # key



# @lc code=end

