#
# @lc app=leetcode id=242 lang=python3
#
# [242] Valid Anagram
#

# @lc code=start
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        d = dict()
        for c in s:
            if not c in d:
                d[c]=0
            d[c] += 1
        for c in t:
            if not c in d:
                return False
            d[c]-=1
        for k,v in d.items():
            if not v==0:
                return False
        return True
        
# @lc code=end

