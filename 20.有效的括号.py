#
# @lc app=leetcode.cn id=20 lang=python3
#
# [20] 有效的括号
#

# @lc code=start
class Solution:
    def isValid(self, s: str) -> bool:
        stack  = []
        for c in s:
            if c in ['(','[','{']:
                stack.append(c)
            elif stack and {"]":"[",")":"(","}":"{"}.get(c) == stack[-1]:
                stack.pop()
            else:
                return False
        return not stack
        

# @lc code=end

if __name__ == "__main__":
    print(Solution().isValid('()'))
    print(Solution().isValid("()[]{}"))
    print(Solution().isValid("(]"))
    print(Solution().isValid("([)]"))