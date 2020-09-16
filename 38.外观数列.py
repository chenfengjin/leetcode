#
# @lc app=leetcode.cn id=38 lang=python3
#
# [38] 外观数列
#

# @lc code=start
class Solution:
    def countAndSay(self, n: int) -> str:
        if n == 1:
            return "1"
        count = 0
        index = 0 
        result = []
        previous = self.countAndSay(n-1) + "#"  #TODO 大一 C语言就学过的技巧，现在终于想起来了 
        while index < len(previous):
            if index and not previous[index] == previous[index-1]:
                result.append(str(count))
                result.append(previous[index-1])
                count = 0
            index += 1
            count += 1
        return "".join(result)
                
# @lc code=end

if __name__ == "__main__":
    for i in range(1,6):
        print(Solution().countAndSay(i))