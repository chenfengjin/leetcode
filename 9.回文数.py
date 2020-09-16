#
# @lc app=leetcode.cn id=9 lang=python3
#
# [9] 回文数
#

# @lc code=start
# TODO 测试用例不完全
class Solution:
    def isPalindrome(self, x: int) -> bool:
        return x>=0 and x==self.reverse(x)
    def reverse(self, x: int) -> int:
        result = 0
        sign  =  1 if  x >=0 else -1
        x = abs(x)
        MAX = (1 << 31)-1 #bug point，貌似有bug？
        # print((MAX))
        while x:
            remain = x %10 
            if (MAX -remain) /10 < +result : # bug point
                return 0
            result *= 10
            result += remain
            # result 
            x //= 10
        return sign * result        
# @lc code=end

