#
# @lc app=leetcode.cn id=5 lang=python3
#
# [5] 最长回文子串
#

# @lc code=start
class Solution:
    # TODO 这种写法比从某一个向两边扩展的方法更好
    # TODO 记录几个典型case ccc abba aba
    # TODO 核心是判断上一个和上上一个之间是不是只差 1
    def longestPalindrome(self, s: str) -> str:
        if len(s) < 2:
            return s
        dp = [1]
        index = 1
        while index < len(s):
            current = 1
            if index > 0 and s[index] == s[index-1]:  #注意这种情况
                current = 2                
            if not index-dp[-1]-1 < 0  and s[index] == s[index-dp[-1]-1]:
                current = max(current,dp[-1]+2)
            index += 1
            dp.append(current)
        index = dp.index(max(dp))
        print(dp)
        return s[index-dp[index]+1:index+1]

# @lc code=end

if __name__ == "__main__":
    print(Solution().longestPalindrome("babad"))
    print(Solution().longestPalindrome("ccc"))