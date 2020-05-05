#
# @lc app=leetcode id=167 lang=python3
#
# [167] Two Sum II - Input array is sorted
#

# @lc code=start
class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        target_map={}
        for index,num in enumerate(numbers):
            if num in target_map.keys():
                break
            target_map[target-num]=index
        index1 = target_map[num]
        index2 = index
        if index1 > index2:
            index1,index2 = index2,index1
        return [index1+1,index2+1]


# @lc code=end

