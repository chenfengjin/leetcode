#
# @lc app=leetcode.cn id=730 lang=python3
#
# [730] 统计不同回文子序列
#

# @lc code=start
class Solution:
    def countPalindromicSubsequences(self, S: str) -> int:
        S = '#' + "#".join([i for i in S]) +'#'
        total = 0
        for index,char in enumerate(S):
            if char == '#':
                continue
            total += 1
            
            
        
# @lc code=end

