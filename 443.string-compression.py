#
# @lc app=leetcode id=443 lang=python3
#
# [443] String Compression
#
import collections
from typing import *
# @lc code=start
class Solution:
    def compress(self, chars: List[str]) -> int:
        if not chars:
            return 0 
        if len(chars) == 1:
            return 1
        left = 0
        right = 1
        count = 1
        while right < len(chars):
            if chars[right] == chars[right-1]:
                right = right + 1
                count  = count + 1
                continue
            else:
                chars[left] =  chars[right-1]
                left += 1
                if not count == 1:
                    chars[left:left+len(str(count))] = [i for i in str(count)]
                    left += len(str(count))   
                count = 1
                right = right + 1 
        chars[left] =  chars[right-1]
        left += 1
        if not count == 1:
            chars[left:left+len(str(count))] = [i for i in str(count)]
            left += len(str(count))
        return left  


# @lc code=end

if __name__ == "__main__":
    print(Solution().compress(["a","a","b","b","c","c","c"]))
    # print(Solution().compress["a"]))

