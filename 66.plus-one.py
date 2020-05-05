#
# @lc app=leetcode id=66 lang=python3
#
# [66] Plus One
#

# @lc code=start
from typing import *
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits.reverse()
        res = []
        carry = 1
        for digit in digits:
            res.append((digit+carry)%10)
            carry = (digit + carry) // 10
        if carry:
            res.append(carry)
        res.reverse()
        return res 
# @lc code=end

if __name__ == "__main__":
    print(Solution().plusOne([9]))
