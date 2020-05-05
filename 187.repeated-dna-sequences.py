#
# @lc app=leetcode id=187 lang=python3
#
# [187] Repeated DNA Sequences
#

# @lc code=start
class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        if len(s) < 10:
            return []
        occur_count = {}
        for i in range(0,len(s)-10 + 1 ) :
            if not  s[i:i+10]  in occur_count:
                occur_count[s[i:i+10]] = 0
            occur_count[s[i:i+10]] += 1
        return [i for i in occur_count if occur_count[i] > 1]
            

        
# @lc code=end