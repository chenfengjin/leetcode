#
# @lc app=leetcode id=73 lang=python3
#
# [73] Set Matrix Zeroes
#

<<<<<<< HEAD
=======
from typing import *
>>>>>>> RFC:change mac
# @lc code=start
class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
<<<<<<< HEAD
        
# @lc code=end

=======
        zeros_positons = []
        rows_count = len(matrix)
        columns_count = len(matrix[0])
        for i in range(rows_count):
            for j in range(columns_count):
                if matrix[i][j] == 0:
                    zeros_positons.append([i,j])
        rows = [i[0] for i in zeros_positons]
        columns = [i[1] for i in zeros_positons]
        for row in rows:
            for column in range(columns_count):
                matrix[row][column] = 0
        for row in range(rows_count):
            for column in columns:
                matrix[row][column] =0            
        
# @lc code=end

if __name__ == "__main__":
    a=[[1,1,1],[1,0,1],[1,1,1]]
    print(a)
    Solution().setZeroes(a)
    print(a)
>>>>>>> RFC:change mac
