#
# @lc app=leetcode.cn id=1550 lang=python3
#
# [1550] 存在连续三个奇数的数组
#
from typing import *
# TODO  有趣的地方在于如何判断奇数
# @lc code=start
class Solution:
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:
        count = 0
        for num in arr:
            count = 0 if not num&1 else count + 1
            if count == 3:
                return True
        return False
# @lc code=end

if __name__ == "__main__":
    print(Solution().threeConsecutiveOdds([2,6,4,1]))
    print(Solution().threeConsecutiveOdds([1,2,34,3,4,5,7,23,12]))
