#
# @lc app=leetcode id=695 lang=python3
#
# [695] Max Area of Island
#

# @lc code=start
class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:
        if not grid or not grid[0]:
            return 0

        islands = []
        rows = len(grid)
        columns = len(grid[0])
        for i in range(n):
            for j in range(n):
                if grid[i][j]:
                    islands.append([i,j])
        max_area = 0
        for island in islands:
            current = island
            while current:
                next_to_search = []
                for row,column in current:
                    grid[row][column] = 0
                    if row + 1 < rows and grid[row+1][column]:
                        next_to_search.append([row+1,column])
                    
                    
        
# @lc code=end

