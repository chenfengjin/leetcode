#
# @lc app=leetcode.cn id=7 lang=python3
#
# [7] 整数反转
#

# @lc code=start
# TODO 好像有问题
class Solution:
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

if __name__ == "__main__":
    print(Solution().reverse(123))