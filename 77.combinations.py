#
# @lc app=leetcode id=77 lang=python
#
# [77] Combinations
#
from typing import *
import copy
# @lc code=start
class Solution(object):
    def __init__(self):
        self.res = []
        self.out = []
    def combine(self, n, k):
        """
        :type n: int
        :type k: int
        :rtype: List[List[int]]
        """
        self.helper(n,k,1)
        return self.res 

    def helper(self,n,k,level):
        if len(self.out) == k:
            self.res.append(copy.copy(self.out))
            return
        for i in range(level,n+1):
            self.out.append(i)
            self.helper(n,k,i+1)
            self.out.pop()
            
# @lc code=end

if __name__ =="__main__":
    print(Solution().combine(n=4,k=2))