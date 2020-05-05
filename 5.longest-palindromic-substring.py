#
# @lc app=leetcode id=5 lang=python3
#
# [5] Longest Palindromic Substring
#

# @lc code=start
class Solution:
    def longestPalindrome(self, s: str) -> str:
        if not s:
            return ""
        if len(s)==1:
            return s
        longest_substring = ""
        for index in range(len(s)):
            substring = self.expand(s,index,index)
            if len(substring) > len(longest_substring):
                    longest_substring = substring
            if not index == len(s)-1 and s[index] == s[index+1]:
                substring = self.expand(s,index,index+1)
                if len(substring) > len(longest_substring):
                    longest_substring = substring
        return longest_substring
                         
    def expand(self,s,left,right):
        if left <0 or right == len(s):
            return s[left+1:right]
        if not s[left] == s[right]:
            return s[left+1:right]
        return self.expand(s,left-1,right+1)
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().longestPalindrome("babad"))