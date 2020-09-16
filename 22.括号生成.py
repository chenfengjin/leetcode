#
# @lc app=leetcode.cn id=22 lang=python3
#
# [22] 括号生成
#
from typing import *
# @lc code=start
# 动态规划的解法
class Solution:
    # def __init__(self):
    def generateParenthesis(self, n: int) -> List[str]:
        result = []
        self.helper(result=result,count=0,current=[],diff=0,n=n)
        return result
    def helper(self,result,count,current,diff,n):
        if diff < 0:
            return 
        if count  == n * 2:
            if not diff:
                result.append("".join(current))
            return
        current.append(")")
        self.helper(result,count+1,current,diff -1,n)
        current.pop()
        current.append("(")
        self.helper(result,count+1,current,diff +1,n)
        current.pop()   
# @lc code=end
# TODO follow up 如果只需要计算括号数呢？

if __name__ == "__main__":
    print(Solution().generateParenthesis(3))