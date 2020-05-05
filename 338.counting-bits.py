#
# @lc app=leetcode id=338 lang=python3
#
# [338] Counting Bits
#

# @lc code=start
class Solution:
    def countBits(self, num: int) -> List[int]:
        # key point: suffix 1 and length of suffix 1 without one and only 1
        suffix_one_count = [0] * (n+1)
        suffix_one_length = [0] * (n+1)
        for i in range(n+1):

        
# @lc code=end

