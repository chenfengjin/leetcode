#
# @lc app=leetcode.cn id=383 lang=python3
#
# [383] 赎金信
#

# @lc code=start
from collections import defaultdict
class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        m = defaultdict(int)
        # 只需要一个 dict 就好了
        for c in magazine:
            m[c] += 1
        for c in ransomNote:
            m[c] -= 1
            if m[c] < 0:
                return False
        return True
            
            

# @lc code=end

