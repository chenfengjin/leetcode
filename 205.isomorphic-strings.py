#
# @lc app=leetcode id=205 lang=python3
#
# [205] Isomorphic Strings
#

# @lc code=start
class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        m1 ={}
        m2 ={}
        for i in range(len(s)):
            if not s[i] in m1:
                m1[s[i]] = t[i]
            if not m1[s[i]] == t[i]:
                return False
            
            if not t[i] in m2:
                m2[t[i]] = s[i]

            if not m2[t[i]] == s[i]:
                return False
        return True
        
# @lc code=end

