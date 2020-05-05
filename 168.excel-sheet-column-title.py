#
# @lc app=leetcode id=168 lang=python3
#
# [168] Excel Sheet Column Title
#

# @lc code=start
class Solution:
    def convertToTitle(self, n: int) -> str:
        n = n - 1
        numbers = []
        while not n//26 == 0: # in this way it will be easy to understand
            numbers.append(n // 26)
            n = n // 26
        numbers.append(n)
        print(numbers)
        
        return "".join([ chr(ord('A') + number - 1) for number in numbers])
    
# @lc code=end

if __name__ == "__main__":
    print(Solution().convertToTitle(26))
    print(Solution().convertToTitle(30))