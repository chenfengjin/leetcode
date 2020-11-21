#
# @lc app=leetcode.cn id=142 lang=python3
#
# [142] 环形链表 II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def detectCycle(self, head: ListNode) -> ListNode:
        l1 = l2 = head
        # leetcode 测试用例没有的情况，只有一个节点指向自身
        if not l2 or not l2.next:
            return None
        # 这里的特殊条件先判断
        while l2 and l2.next:
            l2 = l2.next.next
            l1 = l1.next    
            if l1 == l2:
                break
        if not l1 == l2:
            return None
        
        l1 =  head
        while not l1 == l2:
            l1 = l1.next
            l2 = l2.next    
        return l2
        
        
# @lc code=end

