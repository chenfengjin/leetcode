#
# @lc app=leetcode id=179 lang=python3
#
# [179] Largest Number
#
from typing import *
from functools import cmp_to_key
# @lc code=start
class Solution:
    @staticmethod
    def compare(x,y):
        print(x,y)
        for i in range(min(len(x),len(y))):
            if x[i]<y[i]:
                return -1
            elif x[i] > y[i]:
                return 1
        if len(x) == len(y):
            return 0
        if len(x) < len(y):
            return Solution.compare(y[:i+1],y[i+1:])
        print(i)
        return Solution.compare(x[i+1:],x[:i+1]) # 这里是关键，注意理解两个参数的含义
           
    def largestNumber(self, nums: List[int]) -> str:
        if max(nums) == 0: # important
            return "0"

        key = self.compare
        nums = [str(num) for num in nums]
        nums = sorted(nums,key=cmp_to_key(key),reverse=True)
        return "".join(nums)
        
# @lc code=end
if __name__ == "__main__":
    print(Solution().largestNumber([128,12]))
    # print(Solution().largestNumber([3,30]))
