#
# @lc app=leetcode id=240 lang=python3
#
# [240] Search a 2D Matrix II
#

# @lc code=start
class Solution:
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        if not matrix or not matrix[0]:
            return False
        row_count = len(matrix)
        column_count = len(matrix[0])
        row = 0
        column = column_count - 1
        while row < row_count and not column < 0:
            if matrix[row][column] == target:
                return True
            if matrix[row][column] < target:
                row += 1
                continue
            else:
                column = column - 1
                continue
        return False
        
# @lc code=end

