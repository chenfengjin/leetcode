#
# @lc app=leetcode id=224 lang=python3
#
# [224] Basic Calculator
#

# @lc code=start
import sys


class Solution:
    def calculate(self, s: str) -> int:
        total = [0]
        current = 0
        left = 0
        right = 0
        sign = 1
        s = s.replace(" ", "")
        s = "({})".format(s)
        signs = []

        while right < len(s):

            if s[right] == "(":
                total.append(current)
                current = 0
                left = right
                signs.append(sign)
                sign = 1


            elif s[right] == ")":
                sign = signs.pop()
                current += total.pop() + int("0" + s[left + 1:right]) * sign
                left = right

            elif s[right] == "+":
                current += sign * int("0"+s[left + 1:right])
                sign = 1
                left = right

            elif s[right] == "-":
                current += sign * int("0"+s[left + 1:right])
                sign = -1
                left = right
            right = right + 1

        return current


if __name__ == "__main__":
    print(Solution().calculate("2-(5-6)"))

# @lc code=end
