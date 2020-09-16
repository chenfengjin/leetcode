#
# @lc app=leetcode.cn id=39 lang=python3
#
# [39] 组合总和
#
from typing import *
# @lc code=start
class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        # 这种解法不完全，会有重复
        # if target < 0:
        #     return []
        # result = []
        # if target in candidates:
        #     result.append([target])
        # for candidate in candidates:
        #     result.extend([[candidate] + i for i in self.combinationSum(candidates=candidates,target = target-candidate)])
        # return result
# @lc code=end


if __name__ == "__main__":
    print(Solution().combinationSum(candidates = [2,3,6,7], target = 7))