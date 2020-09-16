#
# @lc app=leetcode id=91 lang=python3
#
# [91] Decode Ways
#

# @lc code=start
class Solution:
    def numDecodings(self, s: str) -> int:
        if not s:
            return 0
        if s.startswith('0'):
            return 0
        if '00' in s.lstrip('0'):
            return 0
        if len(s) == 1:
            return 1
        dp = [1]
        if s[0:2] <= '26' and not s[1] == '0':
            dp.append(2)
        else:
            dp.append(1)
        if s[1]== '0' and s[0:2]> '26':
            return 0

        for index in range(2,len(s)):
            if s[index] == '0':  # 
                if s[index-1]> '2':
                    return 0
                dp.append(dp[-2])
            elif s[index-1] == '0':
                    dp.append(dp[-1])
            elif s[index-1:index+1] <= '26':
                dp.append(dp[-2]+dp[-1])
            else:
                dp.append(dp[-1])
        return dp[-1]
                
        
# @lc code=end

