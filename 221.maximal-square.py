#
# @lc app=leetcode id=221 lang=python3
#
# [221] Maximal Square
#
from typing import *
# @lc code=start
class Solution:
    def maximalSquare(self, matrix: List[List[str]]) -> int:
        if not len(matrix):
            return 0
        row_count = len(matrix)
        column_count = len(matrix[0])
        square = [[0]*column_count for row in range(row_count)]
        for row_index in range(row_count):
            for column_index in range(column_count):
                if matrix[row_index][column_index] == '1':
                    if row_index > 0 and column_index > 0:          
                        square[row_index][column_index] = min(square[row_index-1][column_index-1],square[row_index][column_index-1],square[row_index-1][column_index]) +1
                    else:
                        square[row_index][column_index] = 1

        return max([max(row)for row in square]) ** 2
# @lc code=end

if __name__ == "__main__":
    matrix = [["1","1","1","1","0"],["1","1","1","1","0"],["1","1","1","1","1"],["1","1","1","1","1"],["0","0","1","1","1"]]
    # print(Solution().maximalSquare([["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]))
    print(Solution().maximalSquare(matrix))