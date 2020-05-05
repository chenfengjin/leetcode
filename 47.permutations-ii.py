#
# @lc app=leetcode id=47 lang=python3
#
# [47] Permutations II
#

# @lc code=start
<<<<<<< HEAD
class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        
# @lc code=end

=======
from typing import *
import copy
class Solution:
    def __init__(self):
        self.out = []
        self.res = []
        self.visited = None

    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        self.visited = [False] * len(nums)
        self.helper(nums)
        return self.res

    def helper(self,nums):
        if len(self.out) == len(nums):
            self.res.append(copy.copy(self.out))
            return
        for index,num in enumerate(nums):   
            if self.visited[index]:
                continue
            if index > 0 and num == nums[index-1] and not self.visited[index-1]:
                continue
            self.out.append(num)
            self.visited[index] = True
            self.helper(nums)
            self.out.pop()
            self.visited[index] = False


# @lc code=end


if __name__ == "__main__":
    print(Solution().permuteUnique([1,1,2]))

>>>>>>> RFC:change mac
