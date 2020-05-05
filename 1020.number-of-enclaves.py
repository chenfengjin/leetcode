#
# @lc app=leetcode id=1020 lang=python3
#
# [1020] Number of Enclaves
#

# @lc code=start
class Solution:
    def numEnclaves(self, A: List[List[int]]) -> int:
        if not A or not A[0]:
            return 0
        rows = len(A)
        columns = len(A[0])
        begin_points = []
        for row in range(rows):
            if A[row][0]:
                begin_points.append([row,0])
            if not [row,columns-1]  in begin_points and A[row][columns-1]:
                begin_points.append([row,columns-1])
        for column in range(columns):
            if not [0,column] in begin_points and A[0][column]:
                begin_points.append([0,column])
            if not [rows-1,column] in begin_points and A[rows-1][column]:
                begin_points.append([rows-1,column])
    def search(self,A,point,rows,columns,row,column):
        pass        
# @lc code=end


