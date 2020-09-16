#
# @lc app=leetcode id=88 lang=python3
#
# [88] Merge Sorted Array
#

# @lc code=start
from typing import *
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        index1 = len(nums1) - 1 
        index2 = len(nums2) -1 
        index_total = len(nums1) - 1 
        while  not nums1[index1]:
            index1 -= 1
        while index1 >= 0  and index2 >= 0:
            print("index1:{},index2:{},index_total{}".format(index1,index2,index_total))
            if nums1[index1] > nums2[index2]:
                nums1[index_total] = nums1[index1]
                index1 -= 1
                index_total -= 1
            elif nums1[index1] < nums2[index2]:
                nums1[index_total] = nums2[index2]
                index_total -= 1
                index2 -= 1
            else: 
                nums1[index_total] = nums2[index2]
                index_total -= 1
                index2 -= 1
                index1 -= 1            
        if index2 > 0:
            nums1[:index2] = nums2[:index2]

# @lc code=end


if __name__ == "__main__":
    l1 = [1,2,3,0,0,0]
    l2 = [2,5,6]
    Solution().merge(l1,3,l2,3)
    print(l1)

    