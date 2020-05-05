#
# @lc app=leetcode id=365 lang=python3
#
# [365] Water and Jug Problem
#
import typing
import copy
# @lc code=start
class Solution:
    def canMeasureWater(self, x: int, y: int, z: int) -> bool:
        return z == 0 or (x + y >= z and z % self.gcd(x, y) == 0)
    def gcd(self,x,y):
        return x if y == 0 else self.gcd(y,x%y)
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().canMeasureWater(13,11,1))