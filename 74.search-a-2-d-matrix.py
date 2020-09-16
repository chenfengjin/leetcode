#
# @lc app=leetcode id=74 lang=python3
#
# [74] Search a 2D Matrix
#

# @lc code=start
class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        for row  in matrix:
            for num in row:
                if num == target:
                    return True

        return False
        
# @lc code=end

