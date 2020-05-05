#
# @lc app=leetcode id=260 lang=python3
#
# [260] Single Number III
#
from typing import *
# @lc code=start
class Solution:
    def find_first_1(self,num:int):
        res = 0
        while not (num >>res)&1:
            res+=1
        return res
        
    def singleNumber(self, nums: List[int]) -> List[int]:
        res = 0
        for num in nums:
            res ^= num
        index = self.find_first_1(res)

        group1, group2 = 0, 0
        for num in nums:
            if num & (1<<index):
                group1 ^= num
            else:
                group2 ^= num
        return[group1,group2]
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().find_first_1(7))
