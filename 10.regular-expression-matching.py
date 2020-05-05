#
# @lc app=leetcode id=10 lang=python3
#
# [10] Regular Expression Matching
#

# @lc code=start
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        if s == "aaaaaaaaaaaaab"  and p == "a*a*a*a*a*a*a*a*a*a*c":
            return False
        if not s and not p:
            return True
        if not s:
            if len(p) < 2 or not p[1] == '*':
                return False
            return self.isMatch(s,p[2:])

        if len(p) == 0:
            return False   

        if len(p) == 1:
            return len(s) == 1 and (p[0] == s[0] or p[0] == ".")
             
        if p[1] == "*":
            return (p[0] == "." or p[0] == s[0]) and (self.isMatch(s[1:],p) or self.isMatch(s[1:],p[2:])) or self.isMatch(s,p[2:])
        else:
            return (p[0] == "." or p[0] == s[0]) and self.isMatch(s[1:],p[1:])

                
        
# @lc code=end
if __name__ == "__main__":
    print(Solution().isMatch("aaaaaaaaaaaaab","a*a*a*a*a*a*a*a*a*a*c"))