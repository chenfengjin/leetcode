#
# @lc app=leetcode.cn id=485 lang=python3
#
# [485] 最大连续1的个数
#

# @lc code=start
class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        # return len(max(''.join(map(str,nums)).split('0'))) # TODO 天秀
        nums.insert(0,0)
        nums.append(0)
        count = 0
        max_count = 0
        index = 1
        while index < len(nums):
            if nums[index] and nums[index-1]:
                count += 1
            elif nums[index]:
                count = 1
            else:
                max_count = max(max_count,count)
                count = 0
            index += 1
        nums.pop(0)
        nums.pop()
        return max_count
        # dp = [0]
        # for num in nums:
        #     if num:
        #         dp.append(dp[-1]+1)
        #     else:
        #         dp.append(0)
        # return max(dp)        
# @lc code=end

