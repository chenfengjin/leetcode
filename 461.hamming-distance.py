#
# @lc app=leetcode id=461 lang=python3
#
# [461] Hamming Distance
#

# @lc code=start
class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        z =  x ^ y
        count = 0
        while  z:
            count +=  z % 2
            z =  z // 2
        return count
        
# @lc code=end