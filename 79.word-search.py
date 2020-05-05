#
# @lc app=leetcode id=79 lang=python3
#
# [79] Word Search
#
from typing import *
# @lc code=start
class Solution: #慢，只打败了 50%
    def __init__(self):
        pass
    def exist(self, board: List[List[str]], word: str) -> bool:
        if word == "":
            return True

        if not board or not board[0]:
            return False

        self.board = board
        self.visited = [[False for _ in range(len(board[0]))] for __ in range(len(board))]
        print(self.visited)

        search_points = []
        for row  in range(len(board)):
            for column in range(len(board[0])):
                if board[row][column] == word[0]:
                    search_points.append([row,column])

        for  row,column in search_points:
            if self.helper(board,word,row,column):
                return True
        return False

    def is_valid_position(self,x,y):
        return 0<=x <len(self.board) and 0<= y <len(self.board[0])
        

    def helper(self,board,word,row,column):
        if word == "":
            return True
        if len(word) == 1 and word == board[row][column]:
            return True
        if not board[row][column] == word[0]:
            return False

        self.visited[row][column] = True
        if self.is_valid_position(row,column+1)  and not self.visited[row][column+1] and self.helper(board,word[1:],row,column+1):
            return True
        if self.is_valid_position(row,column-1)  and not self.visited[row][column-1] and self.helper(board,word[1:],row,column-1):
            return True
        if self.is_valid_position(row+1,column)  and not self.visited[row+1][column] and self.helper(board,word[1:],row+1,column):
            return True
        if self.is_valid_position(row-1,column)  and not self.visited[row-1][column] and self.helper(board,word[1:],row-1,column):
            return True
        self.visited[row][column] = False

# @lc code=end

if __name__ == "__main__":
    # print(Solution().exist([["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]],"ABCB"))
    print(Solution().exist([["a"]],"a"))