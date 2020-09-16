#
# @lc app=leetcode id=169 lang=python3
#
# [169] Majority Element
#

# @lc code=start
class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        count_map = {}
        for num in nums:
            if not num in count_map:
                count_map[num] = 0
            count_map[num] += 1
        for k,v in count_map.items():
            if v > len(nums)//2:
                return k       
# @lc code=end

