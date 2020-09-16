#
# @lc app=leetcode id=139 lang=python3
#
# [139] Word Break
#

# @lc code=start
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        if not s:
            return True
        for word in wordDict:
            if s.startswith(word) and self.wordBreak(s[len(word):],wordDict):
                return True
        return False
        
# @lc code=end

