#
# @lc app=leetcode id=8 lang=python3
#
# [8] String to Integer (atoi)
#

# @lc code=start
class Solution:
    def myAtoi(self, str: str) -> int:
        s = str.lstrip(" ")
        sign = 0
        if not s:
            return 0
        if s[0] == '-':
            sign = -1
        if s[0] == '+':
            sign = 1
        start =  0 if not sign else 1
        sign = sign if sign else 1
        end = start
        while end <len(s):
            if not s[end].isnumeric():
                break
            end += 1

        result = 0

        for char in s[start:end]:
            result *= 10
            result += ord(char)-ord('0')
        result = result * sign
        if result < -2147483648:
            return -2147483648
        if result > 2147483647:
            return 2147483647
        return result
        
# @lc code=end

if __name__ == "__main__":
    for s in ["1","01","11"," 1","h","1 2","-1"]:
        print(s,":",Solution().myAtoi(s),end="\n\n")
