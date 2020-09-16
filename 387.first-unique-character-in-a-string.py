#
# @lc app=leetcode id=387 lang=python3
#
# [387] First Unique Character in a String
#

# @lc code=start
from typing import *
class Solution:
    def firstUniqChar(self, s: str) -> int:
        count_map = {}
        for c in s:
            if not c in count_map:
                count_map[c] = 0
            count_map[c] += 1
        for index,c in enumerate(s):
            if count_map[c] == 1:
                return index 
        return -1

        
# @lc code=end

if __name__ == "__main__":
    print(Solution().firstUniqChar("leetcode"))