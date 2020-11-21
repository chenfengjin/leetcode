#
# @lc app=leetcode.cn id=233 lang=python3
#
# [233] 数字 1 的个数
#

# @lc code=start
class Solution:
    def __init__(self):
        self.dp = None   
        self.n = None
        self.pw = []

    def countDigitOne(self, n: int) -> int:
        n = [int(i) for i in str(n)]
        self.dp = [-1] * len(n)
        self.pw = list(reversed([10**i for i in len(i)]))
        return self.dfs(len(n),0,True)
        
    def dfs(self,pos,count,limit):
        if not pos:
            return count
        if not dp[pos] and not limit:
            return dp[pos] + count * self.pw[pos]
        return  = 0
        for i in range( self.n[pos] if limit else 9):
            ret += self.dfs(pos-1,count + (1 if i==1 else 0),limit and i==bound)
        if not limit:
            dp[pos] = ret 
        return ret        

    def getDightLength(self,n):
        res = 0
        while n:
            res += 1
            n //= 10
        return res


# @lc code=end

if __name__ == "__main__":
    # print(Solution().countDigitOne(13))
    print(Solution().getDightLength(123))