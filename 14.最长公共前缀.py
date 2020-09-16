#
# @lc app=leetcode.cn id=14 lang=python3
#
# [14] 最长公共前缀
#
from typing import *
# @lc code=start
#TODO python break 只停止最深一层
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ""
        min_length = min([len(s) for s in strs])
        found = False
        i = 0
        while i < min_length:
            for s in strs:
                if  found  or  not s[i] == strs[0][i]:
                    found = True
            if found:
                break   
            i+=1
        return strs[0][:i]

# @lc code=end

if __name__ == "__main__":
    print(Solution().longestCommonPrefix(["flower","flow","flight"]))