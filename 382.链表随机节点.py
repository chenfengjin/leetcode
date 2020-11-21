#
# @lc app=leetcode.cn id=382 lang=python3
#
# [382] 链表随机节点
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:

    def __init__(self, head: ListNode):
        """
        @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node.
        """
        self.head = head   
        

    def getRandom(self) -> int:
        """
        Returns a random node's value.
        """
        n = 1
        result = self.head  

        current = self.head
        while current.next:
            current = current.next
            if random.randint(1,n+1) == 1:
                result = current
            n += 1
        return result.val



# Your Solution object will be instantiated and called as such:
# obj = Solution(head)
# param_1 = obj.getRandom()
# @lc code=end

