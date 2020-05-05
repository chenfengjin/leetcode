#
# @lc app=leetcode id=125 lang=python3
#
# [125] Valid Palindrome
#

# @lc code=start
class Solution:
    def isPalindrome(self, s: str) -> bool:
        if len(s) <= 1:
            return True
        s = s.lower()
        s=[c for c in s if c.isalnum()]
        length = len(s)
        i = 0
        while i < length -1 -i:
            if not s[i] == s[length-1-i]:
                return False
            i = i + 1
        return True

        
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().isPalindrome("A man, a plan, a canal: Panama"))

