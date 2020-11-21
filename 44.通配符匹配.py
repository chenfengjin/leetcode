#
# @lc app=leetcode.cn id=44 lang=python3
#
# [44] 通配符匹配
#

# @lc code=start
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        if not s or not p:
            return False
        dp = [[False] * len(p)]* len(s)
        # len(s) 行
        # len(p) 列
        dp[0][0] = s[0]==p[0] or p[0] in ['*',"?"]
        for i in range(len(s)):
            dp[i][0] == True if p[0] == '*' or p[0] == '+' or p[0] == s[i] else False
        for i in range(len(p)):
            dp[0][i] = p[i] == '*' or p[i] == '+' or p[i] == s[0]
        for i in range(1,len(s)):
            for j in range(1,len(p)):
                if dp[i-1][j-1] and p[i] in ['*','+'] or s[i] == p[j]:
                    dp[i][j] = True
                elif True:
                    pass
        return dp[-1][-1]


# @lc code=end

if __name__ == "__main__":
    print(Solution().isMatch("aa","a"))