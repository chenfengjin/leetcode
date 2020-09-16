#
# @lc app=leetcode.cn id=201 lang=python3
#
# [201] 数字范围按位与
#

# @lc code=start
class Solution:
    def rangeBitwiseAnd(self, m: int, n: int) -> int:
        # AC but slow
        # i = 31
        # while i >= 0:
        #     if ((m>>i)&1) ^ ((n>>i) &1):
        #         break
        #     i -= 1
        # while i>=0:
        #     m &= (-1)<<i
        #     i -= 1
        # return m
        # still slow
        # i = 31
        # while i>=0  and not ((m>>i)&1) ^ ((n>>i) &1):
        #     i -= 1
        # return m&(-1) << max(0,i)
        while m<n:
            n&=(n-1)
        return n   
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().rangeBitwiseAnd(5,7))