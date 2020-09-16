#
# @lc app=leetcode.cn id=77 lang=python3
#
# [77] 组合
#
from typing import *
import copy
# @lc code=start
class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        result = []
        current = []
        self.helper(result,1,current,n,k)
        return result
    def helper(self,result,index,current,n,k):
        if len(current) == k:
            result.append(copy.deepcopy(current))
            return
        for i in range(index,n +1 ):
            current.append(i)
            self.helper(result,i+1,current,n,k)
            current.pop()

# @lc code=end


if __name__ == "__main__":
    print(Solution().combine(4,2))