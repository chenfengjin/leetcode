#
# @lc app=leetcode id=1 lang=python3
#
# [1] Two Sum
#

# @lc code=start
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        target_map={}
        for index,num in enumerate(nums):
            if num in target_map.keys():
                return [target_map[num],index]
            target_map[target-num]=index

# @lc code=end

