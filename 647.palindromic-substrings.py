#
# @lc app=leetcode id=647 lang=python3
#
# [647] Palindromic Substrings
#

# @lc code=start
class Solution:
    def __init__(self):
        self.count = 0
    
    def countSubstrings(self, s: str) -> int:
        for i in range(len(s)):
            self.count += 1
            self.expand(s,i,i)
        for i in range(len(s)-1):
            if s[i] == s[i+1]:
                self.count += 1
                self.expand(s,i,i+1)
        return self.count    
    def expand(self,s,i,j):
        if i == 0:
            return
        if j == len(s) - 1:
            return 
        if s[i-1] == s[j+1]:
            self.count += 1
            self.expand(s,i-1,j+1)
        
# @lc code=end