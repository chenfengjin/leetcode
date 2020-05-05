#
# @lc app=leetcode id=39 lang=python3
#
# [39] Combination Sum
#

# @lc code=start
import copy
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        result = []
        choices = []
        for index in range(   [-1],len(candidates)):
            if sum([candidates[i] for i in choices]) == target:
                result.append([candidates[i] for i in choices])
            elif sum([candidates[i] for i in choices]) < target:
                choices.append(choices[-1])
            elif sum([candidates[i] for i in choices]) > target:
                choices.pop()
                choices.pop()  
                choices.append(choices[-1]+1)             
# @lc code=end
