#
# @lc app=leetcode id=54 lang=python3
#
# [54] Spiral Matrix
#

# @lc code=start
class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        num = 0
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

