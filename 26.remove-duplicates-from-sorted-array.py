#
# @lc app=leetcode id=26 lang=python3
#
# [26] Remove Duplicates from Sorted Array
#
# @lc code=start
class Solution:
    def removeDuplicates(self, nums):
    # def removeDuplicates(self, nums: List[int]) -> int:
        index = len(nums)-1
        if not len(nums):
            return 0
        if not len(nums)-1:
            return 1
        while index > 0:
            if nums[index]==nums[index-1]:
                nums.remove(nums[index])
            index=index-1
        return len(nums)
# @lc code=end

if __name__ == "__main__":
    print(Solution().removeDuplicates([0,0,1,1,1,2,2,3,3,4],))