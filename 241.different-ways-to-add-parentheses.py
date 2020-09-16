#
# @lc app=leetcode id=241 lang=python3
#
# [241] Different Ways to Add Parentheses
#

# @lc code=start
class Solution:
    def __init__(self):
        self.cache={}
    def diffWaysToCompute(self, s: str) -> List[int]:
        if s in self.cache:
            return self.cache.get(s)

        if s.count("+") + s.count("-") + s.count("*") <= 1:
            return [eval(s)]
        result = []
        for index,c in enumerate(s):
            if c in ['+','-','*']:
                left_values = self.diffWaysToCompute(s[0:index])
                right_values = self.diffWaysToCompute(s[index + 1:])
                self.cache[s[0:index]]=left_values
                self.cache[s[index+1:]]=right_values
                for left_value in left_values:
                    for right_value in right_values:
                        result.append(
                            {
                                "+":left_value + right_value,
                                "-": left_value - right_value,
                                "*" :left_value * right_value,
                            }.get(c)
                        )
        # self.cache[s]=result
        return result
                        
                
                   
# @lc code=end

