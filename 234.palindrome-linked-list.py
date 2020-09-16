#
# @lc app=leetcode id=234 lang=python3
#
# [234] Palindrome Linked List
#

# @lc code=start
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def isPalindrome(self, head: ListNode) -> bool:
        res = []
        if not head:
            return True
        while head:
            res.append(head.val)
            head = head.next
        print(res)
        return res == list(reversed(res))
        
        
# @lc code=end

