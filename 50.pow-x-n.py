#
# @lc app=leetcode.cn id=50 lang=python3
#
# [50] Pow(x, n)
#
# 快速幂算法
# @lc code=start
class Solution:
    def myPow(self, x: float, n: int) -> float:
        dp = [x]
        count = 1
        total = 1
        while total + count <= n or count > 1:
            if count > 1:
                count  = count // 2
            total *= dp[-1]
            dp.append(dp[-1]*dp[-1])
# @lc code=end

