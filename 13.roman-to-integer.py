#
# @lc app=leetcode id=13 lang=python3
#
# [13] Roman to Integer
#

# @lc code=start
<<<<<<< HEAD
class Solution:
    def romanToInt(self, s: str) -> int:
        
# @lc code=end

=======


class Solution:
    def romanToInt(self, s: str) -> int:
        m = {
            "I": 1,
            "V": 5,
            "X":10,
            "L":50,
            "C":100,
            "D":500,
            "M":1000,
        }
        current = 0
        for c in s:
            num = m.get(c)
            if num > current:
                current = num - current
            else:
                current  = num + current  
        return current
# @lc code=end

if __name__ == "__main__":
    # print(Solution().romanToInt('V'))
    # print(Solution().romanToInt('VI'))
    # print(Solution().romanToInt('IV'))
    print(Solution().romanToInt('MCMXCIV'))
>>>>>>> RFC:change mac
