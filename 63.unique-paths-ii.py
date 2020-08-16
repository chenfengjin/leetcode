#
# @lc app=leetcode id=63 lang=python3
#
# [63] Unique Paths II
#
from typing import *

# @lc code=start
class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        if not len(obstacleGrid):
            return 0
        path_count = [[0]*len(obstacleGrid[0]) for _ in obstacleGrid]
        path_count[0][0] = 1
        for row_index in range(len(obstacleGrid)):
            for column_index in range(len(obstacleGrid[0])):
                if not row_index and not column_index and not obstacleGrid[row_index][column_index]:
                    path_count[row_index][column_index] = 1
                    continue
                if obstacleGrid[row_index][column_index]:
                    path_count[row_index][column_index] = 0
                else: 
                    if row_index >0 and column_index >0:
                        path_count[row_index][column_index] =  path_count[row_index][column_index-1] + \
                        path_count[row_index-1][column_index]
                    elif row_index > 0:
                         path_count[row_index][column_index] = path_count[row_index-1][column_index]
                    else:
                        path_count[row_index][column_index] = path_count[row_index][column_index-1]
        return  path_count[-1][-1]

# @lc code=end


if __name__ == "__main__":
    print(Solution().uniquePathsWithObstacles([[0,0,0],[0,1,0],[0,0,0]]))
