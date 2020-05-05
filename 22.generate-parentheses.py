#
# @lc app=leetcode id=22 lang=python3
#
# [22] Generate Parentheses
#
from typing import *
import copy
# @lc code=start
class Solution:
    def __init__(self):
        self.left = 0
        self.right = 0
        self.parentheses = []
        self.parenthese = []
    def generateParenthesis2(self, n: int) -> List[str]:
        n = n * 2
        self.helper(n,0)  
        return self.parentheses
        
    def helper(self,n,k):
        if self.right > self.left:
            return 
        if n==k:
            if self.left == self.right :
                self.parentheses.append("".join(self.parenthese))
            return
        self.parenthese.append("(")
        self.left += 1
        self.helper(n,k+1)
        self.left -= 1
        self.parenthese.pop()

        self.parenthese.append(")")
        self.right += 1
        self.helper(n,k+1)
        self.right -= 1
        self.parenthese.pop()

    def generateParenthesis(self, n: int) -> List[str]:
        parenthesis = [''] * (n * 2)
        left = 0
        right = 0
        
        for s in range(['(',')']):
            for i in range(n*2):

                        
# @lc code=end

if __name__ == "__main__":
    print(Solution().generateParenthesis(3))