#
# @lc app=leetcode id=89 lang=python3
#
# [89] Gray Code
#
from typing import *
# @lc code=start
class Solution:
    def grayCode(self, n: int) -> List[int]:
        if n == 0:
            return [0]
        if n == 1:
            return [0,1]
        result = self.grayCode(n-1)
        right = [i + (1 << (n-1)) for i in reversed(result)]
        result.extend(right)
        return result

# @lc code=end


if __name__ == "__main__":
    print(Solution().grayCode(2))
