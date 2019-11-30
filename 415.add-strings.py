#
# @lc app=leetcode id=415 lang=python3
#
# [415] Add Strings
#

# @lc code=start
class Solution:
    def addStrings(self, num1: str, num2: str) -> str:
        length = max(len(num1),len(num2))
        num2=num2.zfill(length)[::-1]
        num1=num1.zfill(length)[::-1]
        result = ['0']*length
        carry = 0
        for i in range(length):
            result[i] = (ord(num1[i])+ord(num2[i])+carry-ord('0')*2) % 10 +ord('0')
            carry =(ord(num1[i])+ord(num2[i])+carry-ord('0')*2) // 10
        if carry:
            result.append(ord("1"))
        result=result[::-1]
        return "".join([chr(r) for r in result])

# @lc code=end



if __name__ == "__main__":
    print(Solution().addStrings("0","0"))