#
# @lc app=leetcode id=150 lang=python3
#
# [150] Evaluate Reverse Polish Notation
#
from typing import *

# @lc code=start
class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        # if len(tokens) == 1:
        #     return int(tokens[0])
        stack = []
        for token in tokens:
            if token in ["+","-","*","/"]:
                right = stack.pop()
                left = stack.pop()
                value =  {
                        "+":left+right,
                        "-":left-right,
                        "*":left * right,
                        "/": 0 if right == 0 else left//right +1 if left * right < 0  
                        and not left % right == 0 else left//right,
                    }.get(token)
                print(value)
                stack.append(
                    value
                )
            else:
                stack.append(int(token))
        return stack[0]                     
# @lc code=end


if __name__ == "__main__":
    print(Solution().evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"]))
