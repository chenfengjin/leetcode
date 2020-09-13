#
# @lc app=leetcode id=18 lang=python3
#
# [18] 4Sum
#
from typing import *
# @lc code=start
class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        result = []
        l = len(nums)
        for i in range(l):
            for j in range(i+1,l):
                map = {}
                for k in range(j+1,l):
                    if nums[k] in map:
                        m = map[nums[k]]
                        if not sorted([nums[i],nums[j],nums[k],nums[m]]) in result:
                            result.append(sorted([nums[i],nums[j],nums[k],nums[m]]))
                    else:
                        map[target-nums[i]-nums[j]-nums[k] ] = nums[k]
                        
                    # for m in range(k+1,l):
                    #     if nums[i] + nums[j] + nums[k] + nums[m] == target:
        return result
# @lc code=end

if __name__ == "__main__":
    print(Solution().fourSum([1,0,-1,0,-2,2],0))