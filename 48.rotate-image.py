#
# @lc app=leetcode id=48 lang=python3
#
# [48] Rotate Image
#
from typing import *
# @lc code=start
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """

        n = len(matrix)-1
        for row in range(len(matrix)):
            for column in range(row, n-row):
                matrix[row][column],  \
                matrix[n-column][row],\
                matrix[n-row][n-column],\
                matrix[column][n-row] = matrix[n-column][row],\
                matrix[n-row][n-column],\
                matrix[column][n-row] ,\
                matrix[row][column]
            
# @lc code=end
# n=3    row = 0 column = 2

if __name__ == "__main__":
    Solution().rotate([[1,2,3],[4,5,6],[7,8,9]])
# 0 2 -> 0    2
# 1 0 ->  3-2  0
# 3 1 -> 3-0 3-2
# 2 3 -> 2 3-0 

#  1  2  3  4 
#  5  6  7  8 
#  9 10 11 12
# 13 14 15 16 

