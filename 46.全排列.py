#
# @lc app=leetcode.cn id=46 lang=python3
#
# [46] 全排列
#
from typing import *
import copy
# @lc code=start
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        result = []
        visited = [False]*len(nums)
        self.helper(result,nums,[],visited)
        return result
    def helper(self,result,nums:List[int],current,visited):
        if  len(current)==len(nums):
            result.append(current[:])
            return
        for index,num  in enumerate(nums):
            if visited[index]:
                continue
            visited[index] = True
            current.append(num)
            self.helper(result,nums,current,visited)
            current.pop()
            visited[index] = False


# @lc code=end

if __name__ == "__main__":
    print(Solution().permute([1,2,3]))