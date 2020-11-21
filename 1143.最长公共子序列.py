#
# @lc app=leetcode.cn id=1143 lang=python3
#
# [1143] 最长公共子序列
#
from typing import *
# @lc code=start
class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        if not text1 or not text2:
            return 0
        # dp = [[0]*len(text1)]* len(text2) # T1 列，T2 行
        dp = [[0 for _ in text1] for _ in text2]
        dp[0][0] =  1 if text1[0] == text2[0] else 0
        print(id(dp[0]))
        print(id(dp[1]))
        for i in range(1,len(text1)):
            dp[0][i] = max(1 if  text2[0] == text1[i] else 0,dp[0][0])
        for i in range(1,len(text2)):
            dp[i][0] = max(1 if text1[0] == text2[i] else 0,dp[0[0]])
            
        for index1,char1 in enumerate(text2[1:],start = 1):
            for index2,char2 in enumerate(text1[1:],start = 1):
                if char1 == char2:
                    dp[index1][index2] = dp[index1-1][index2-1] + 1 
                else:
                    dp[index1][index2] = max(dp[index1-1][index2],dp[index1][index2-1])
        
        return max([max([i]) for i in dp])
# @lc code=end


if __name__ == "__main__":
    print(Solution().longestCommonSubsequence("abcde","ace"))