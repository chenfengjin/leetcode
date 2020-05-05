#
# @lc app=leetcode id=400 lang=python3
#
# [400] Nth Digit
#

# @lc code=start
from typing import *
class Solution:
    def findNthDigit(self, n: int) -> int:
        if n < 9:
            return n
        range_map = [9]
        for i in range(1,16):
            last = range_map[len(range_map)-1]
            range_map.append(last+(i+1)*9*(10**i))
        for index,r in enumerate(range_map):
            if r > n:
                break
        remains = n-range_map[index-1] 
        count = remains // (index+1) 
        bit = remains %(index+1)
        count = count +1 if  bit else count
        s = str(10**index-1+count)
        return s[bit-1]
        
# @lc code=end

if __name__ =="__main__":
    print(Solution().findNthDigit(1000))


