#
# @lc app=leetcode id=67 lang=python3
#
# [67] Add Binary
#

# @lc code=start
class Solution:
    # def char_to_int(ch)
    def addBinary(self, a: str, b: str) -> str:
        if len(a)>len(b):
            a,b = b, a
        a = '0' * (len(b) - len(a)) + a
        index  = len(a) - 1
        carry = 0
        result = []
        while index >= 0:
            sum = carry + ord(a[index]) -ord('0') + ord(b[index]) - ord('0')
            carry = sum//2
            result.append( sum %2)
            index -= 1
        if carry:
            result.append(1)
        # print(result)
        
        return "".join([chr(i+ord('0')) for i in reversed(result)])     
        
# @lc code=end

if __name__ == "__main__":
    Solution().addBinary("001","00001")
    