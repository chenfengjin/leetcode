#
# @lc app=leetcode id=59 lang=python3
#
# [59] Spiral Matrix II
#

# @lc code=start
from typing import *
class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        num = 1
        # directions = [0,1,2,3] # right,down,left up
        direction = 0
        matrix = [[0 for i in range(n)] for j in range(n)]
        row = 0
        column = 0
        while num < n*n + 1:
            matrix[row][column] = num
            num += 1
            if direction == 0:
                if not column +1 < n or matrix[column+1][row]:
                    direction  = 1   
                    row += 1
                else:
                    column += 1
                continue
                    
            if direction == 1:
                if not row +1 < n or matrix[row+1][column]:
                    direction = 2
                    column -= 1
                else:
                    row +=1 
                continue
            if direction == 2:
                if column -1 < 0 or matrix[row][column-1]:
                    direction = 3
                    row -= 1
                else:
                    column -=1
                continue
            if direction == 3:
                if row -1 < 0 or matrix[row-1][column]:
                    direction = 0
                    column += 1
                else:
                    row -= 1
                continue
        return matrix
       
# @lc code=end


if __name__ == "__main__":
    matric = Solution().generateMatrix(4)
    for row in matric:
        print(row)