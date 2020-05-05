#
# @lc app=leetcode id=62 lang=python3
#
# [62] Unique Paths
#

# @lc code=start
class Solution:
    def combination(self,n,k): # 坑在浮点预算
        res = 1
        i = k
        while i > 0:
            res *=  n
            n = n-1
            i -= 1
        i = k
        while i > 0 :
            res = res // i
            i -= 1
        return res
    def uniquePaths(self, m: int, n: int) -> int:
        return self.combination(m+n-2,n-1)
        
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().combination(100,1))