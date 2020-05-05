#
# @lc app=leetcode id=118 lang=python3
#
# [118] Pascal's Triangle
#

# @lc code=start
class Solution:
    def factorial(self,n):
        return 1 if n <=  1 else self.factorial(n-1) * n
    def generate(self, numRows: int) -> List[List[int]]:
        res = []
        for row_index in range(numRows):
            row = []
            for column_index in range(row_index):
                row.append(self.factorial(row_index-1)/self.factorial(row_index-1)/self.factorial(row_index-column_index))

        # (n-1)!/[(m-1)!(n-m)!]
        
# @lc code=end

