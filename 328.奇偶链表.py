#
# @lc app=leetcode.cn id=328 lang=python3
#
# [328] 奇偶链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def oddEvenList(self, head: ListNode) -> ListNode:
        flag = 0
        odd = ListNode()
        odd_head = odd
        even = ListNode()
        even_head = even

        while head:
            if flag:
                odd.next = head
                odd = odd.next
            else:
                even.next = head
                even = even.next
            flag = not flag
            head = head.next
        odd.next = None
        even.next = odd_head.next
    
        return even_head.next 
              
# @lc code=end

