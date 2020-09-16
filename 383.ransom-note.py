#
# @lc app=leetcode id=383 lang=python3
#
# [383] Ransom Note
#

# @lc code=start
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        word_count = {}
        for c in magazine:
            if not c in word_count:
                word_count[c] = 0
            word_count[c] += 1
        for c in ransomNote:
            if not c in word_count or not word_count[c]:
                return False
            word_count[c] -= 1
        return True
        
# @lc code=end

