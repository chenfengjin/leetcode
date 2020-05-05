#
# @lc app=leetcode id=119 lang=python3
#
# [119] Pascal's Triangle II
#
from typing import *
# @lc code=start
class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        current = 1
        result = [1]
        for i in range(rowIndex+1):
            current = current * (rowIndex -i +1) // i
            result.append(current)

        return result



        
# @lc code=end

if __name__ == "__main__":
    s=Solution()
    for i in  range(4):
        print(s.combanation(4,i))