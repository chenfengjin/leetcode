#
# @lc app=leetcode id=118 lang=python3
#
# [118] Pascal's Triangle
#
from typing import *
import copy

# @lc code=start
class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        result = []
        for i in range(numRows):
            res = [1]
            current = 1
            for j in range(i):
                current = current * (i - j) // (j+1)
                res.append(current)      
            result.append(res)
        return result        
# @lc code=end

if __name__ == "__main__":
    print(Solution().generate(5))
