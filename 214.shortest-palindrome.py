#
# @lc app=leetcode id=214 lang=python3
#
# [214] Shortest Palindrome
#

# @lc code=start
class Solution: # 本质上是找最长回文前缀
    def is_palindrome(self,s):
        if len(s) == 0:
            return True
        left = 0
        right = len(s) - 1
        while left <= right:
            if not s[left] == s[right]:
                return False
            left += 1
            right -= 1
        return True
    def shortestPalindrome(self, s: str) -> str:
        if not s:
            return ""
        right = len(s)
        s = [i for i in s]
        while right > 0:
            if self.is_palindrome(s[:right]):
                return "".join(s[right:][::-1] + s)
            else:
                # print(right,self.is_palindrome(s[:right]))
                pass
            right -= 1
        
# @lc code=end

if __name__ == "__main__":
    # print(Solution().is_palindrome("a"))
    # print(Solution().is_palindrome("aba"))
    # print(Solution().is_palindrome("abba"))
    s=""
    # print(len(s))
    print(Solution().shortestPalindrome(s))
