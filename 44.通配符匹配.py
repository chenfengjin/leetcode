#
# @lc app=leetcode.cn id=44 lang=python3
#
# [44] 通配符匹配
#

# @lc code=start
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        while "**" in p:
            p.replace("**","*")
        index_s= 0
        index_p = 0
        while index_s  < len(s) and index_p < len(p):
            if s[index_s] == p[index_p]:
                index_p += 1
                index_s += 1
            elif p[index_p] == '?':
                index_p += 1
                index_s += 1
            elif p[index_p] == '*':
                index_s = len(index_s) - 1
            else:
                return False
        # return index_p == len(p) and index_s == len(s)
        





# @lc code=end

