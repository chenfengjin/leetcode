#
# @lc app=leetcode id=282 lang=python3
#
# [282] Expression Add Operators
#
from typing import *
# @lc code=start
class Solution:
    def __init__(self):
        self.res = []
    def addOperators(self, num: str, target: int) -> List[str]:
        print(num,target)
        if not num:
            return []

        if int(num) == target:
            return num

        if len(num) == 1 and not int(num) ==  target :
            return []
        result = []
        num_left = int(num[0])
        num_right = num[1:]
        result.extend(
            [ 
                "{}+{}".format(num_left,i) 
                for i in self.addOperators( num_right,target - num_left) 
            ])
        result.extend(
            [ 
                "{}-{}".format(num_left,i) 
                for i in self.addOperators( num_right,target + num_left) 
            ])
        if target % num_left == 0:
            result.extend(
            [ 
                "{}*{}".format(num_left,i) 
                for i in self.addOperators( num_right,target // num_left) 
            ])
        return result
        
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().addOperators("123",6))