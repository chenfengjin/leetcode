#
# @lc app=leetcode.cn id=342 lang=python3
#
# [342] 4的幂
#

# @lc code=start
# 注意负数都不对
# 注意几个边界值 0 1
class Solution:
    def isPowerOfFour(self, num: int) -> bool:
        return num in [
                1,
                1<<2,
                1<<4,
                1<<6,
                1<<8,
                1<<10,
                1<<12,
                1<<14,
                1<<16,
                1<<18,
                1<<20,
                1<<22,
                1<<24,
                1<<26,
                1<<28,
                1<<30,
        ]
        
# @lc code=end

