#
# @lc app=leetcode id=394 lang=python3
#
# [394] Decode String
#
from typing import *
# @lc code=start
class Solution:
    LETTER = 0
    NUMBER = 1
    BRACKET = 2
    def get_type(self,c:str):
        if c.isalpha():
            return self.LETTER
        if c.isnumeric():
            return self.NUMBER
        if c in ["[","]"]:
            return self.BRACKET
    def decodeString(self, s: str) -> str:
        brackets = []
        letters = []
        numbers =[]
        left = 0
        inner = ""
        for right in range(len(s)):
            if self.get_type(s[left]) == self.get_type(right):
                continue
            if self.get_type(s[left]) == self.LETTER:
                letters.append(s[left:right])
            if self.get_type(s[left]) == self.NUMBER:
                numbers.append(int(s[left:right]))
            if self.get_type(s[left]) == self.BRACKET:
                pass
            right = left
# @lc code=end

if __name__ == "__main__":
    for c in "i2[]" :
        print(Solution().get_type(c))

