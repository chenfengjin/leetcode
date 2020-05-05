#
# @lc app=leetcode id=88 lang=python3
#
# [88] Merge Sorted Array
#

# @lc code=start
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        index1 = 0
        index2 = 0
        while nums1[index1] >= nums1[index1-1]:
            if nums1[index1] >= nums2[index2]:
                index1 += 1
                continue
            else:
                num1[index1],nums2[index2] = nums2[index2],nums1[index1]
                index1 += 1
# @lc code=end


if __name__ == "__main__":
    l1 = [1,3,5,7,9,0,0,0]
    l2 = [2,4,8]
    Solution().merge(l1,l2)
    print(l1)

    