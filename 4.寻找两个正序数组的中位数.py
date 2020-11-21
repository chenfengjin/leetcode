#
# @lc app=leetcode.cn id=4 lang=python3
#
# [4] 寻找两个正序数组的中位数
#

# @lc code=start
class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        l1,l2 = 0
        r1,r2 = len(nums1),len(nums2)

        m1 = (l1+r1)//2
        m2 = (l2+r2)//2

        while m1 + m2 < (len(nums1) +len(nums2))//2:
            if nums1[m1] > nums2[m2]:
                pass
            elif nums1[m1] < nums2[m2]:
                pass
            else:
                pass

# @lc code=end

