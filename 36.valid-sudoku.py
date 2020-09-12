#
# @lc app=leetcode id=36 lang=python3
#
# [36] Valid Sudoku
#

# @lc code=start
from typing import *
class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        for row in board:
            if not self.subValid([row]):
                return False
        for i in range(0,9):
            if not self.subValid([row[i] for row in board]):
                return False
        
        for i in [0,3,6]:
            for j in [0,3,6]:
                if not self.subValid([column[j:j+3] for column in board[i:i+3]]):
                    return False
        return True
        
    def subValid(self,board):
        m = set()
        for row in board:
            for ele in row:
                if ele == '.':
                    continue
                if ele in m:
                    return False
                m.add(ele)
        return True
# @lc code=end
if __name__ == "__main__":
    board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
    print(Solution().isValidSudoku(board))

