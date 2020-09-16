#
# @lc app=leetcode id=38 lang=python3
#
# [38] Count and Say
#

# @lc code=start
class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"
        previous = self.countAndSay(n-1)
        result = []
        c1 = None # key 
        count = 0
        for c in previous:
            if not c1:  # handle the first character
                c1 = c    
            if not c == c1:
                result.append(str(count))
                result.append(c1)
                count = 1
                c1 = c
            else:
                count += 1
        result.append(str(count))
        result.append(c)
        return "".join(result)
                                    
# @lc code=end


if __name__ == "__main__":
    for i in range(1,10):
        print(i,Solution().countAndSay(i),sep=':')
