#
# @lc app=leetcode id=119 lang=python3
#
# [119] Pascal's Triangle II
#
from typing import *
# @lc code=start
class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        res = [1]
        current = 1
        for j in range(rowIndex):
            current = current * (rowIndex - j) // (j+1)
            res.append(current)      
        return res

        
# @lc code=end
