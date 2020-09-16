#
# @lc app=leetcode.cn id=47 lang=python3
#
# [47] 全排列 II
#
from typing import *
# @lc code=start
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        result = []
        visited = [False]*len(nums)
        self.helper(result,nums,[],visited)
        return result
    def helper(self,result,nums:List[int],current,visited):
        if  len(current)==len(nums):
            result.append(current[:])
            return
        for index,num  in enumerate(nums):
            if visited[index] :
                continue
            # TODO 这一行好难写
            # TODO 就加了这三行
            elements = [flag for i,flag in enumerate(visited) if i < index and nums[i] == nums[index] ]
            if elements and  not all(elements):
                continue            
            visited[index] = True
            current.append(num)
            self.helper(result,nums,current,visited)
            current.pop()
            visited[index] = False
# @lc code=end

if __name__ == "__main__":
    print(Solution().permuteUnique([1,1,2]))
