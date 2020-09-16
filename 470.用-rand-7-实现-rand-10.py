#
# @lc app=leetcode.cn id=470 lang=python3
#
# [470] 用 Rand7() 实现 Rand10()
#

# @lc code=start
# The rand7() API is already defined for you.
# def rand7():
# @return a random integer in the range 1 to 7

class Solution:
    def rand10(self):
        """
        :rtype: int
        """
        a = rand7()
        b = rand7()
        while a + (b-1)*7 > 40:
            a = rand7()
            b = rand7()
        # return (a*b //4) +1
        return 1 + (a + (b-1)*7 - 1)   % 10
        
# @lc code=end

