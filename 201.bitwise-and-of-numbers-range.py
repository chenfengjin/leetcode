#
# @lc app=leetcode id=201 lang=python3
#
# [201] Bitwise AND of Numbers Range
#
from typing import *
# @lc code=start
class Solution:
    def rangeBitwiseAnd(self, m: int, n: int) -> int:
    # 找到最小的值的第一个 1，然后把最大值这个位置后边的都编程0 就okla
    # 所以需要一个变量来记录位置？ No!!
    #  101
    #  100 
        if m == 0:
            return 0
        if m == n:
            return n
        if m == 1:
            return m & -2  

        i = 1     
        while m>>1:
            m >>= 1
            n &= (-1)<<(i)
            i += 1
        return n

# @lc code=end

#  101
#  100

# 0100
# 0010


if __name__ == "__main__":
    print(Solution().rangeBitwiseAnd(5,7))
    print(Solution().rangeBitwiseAnd(1,2))