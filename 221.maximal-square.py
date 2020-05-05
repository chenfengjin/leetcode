#
# @lc app=leetcode id=221 lang=python3
#
# [221] Maximal Square
#
from typing import *
# @lc code=start
class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        if not matrix or not matrix[0]:
            return 0
        search_points = []
        row_count = len(matrix)
        column_count = len(matrix[0])
        for row in range(row_count):
            for column in range(column_count):
                if not matrix[row][column]=="1":
                    continue
                if row == 0 and column == 0:
                    search_points.append([row,column])
                    continue
                if not row ==0:
                    
        for row in matrix:
            print(row)
        print(search_points)
# @lc code=end

if __name__ == "__main__":
    print(Solution().maximalSquare([["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]))