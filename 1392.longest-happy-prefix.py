#
# @lc app=leetcode id=1392 lang=python3
#
# [1392] Longest Happy Prefix
#

# @lc code=start
class Solution:
    def longestPrefix(self, s: str) -> str:
        i = 0
        while s[:i+1] ==s[-i-1:] and i < len(s):
            i+=1
        if i == 0:
            return ""
        return s[:i]
        

        
# @lc code=end

