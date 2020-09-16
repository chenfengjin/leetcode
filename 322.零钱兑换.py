#
# @lc app=leetcode.cn id=322 lang=python3
#
# [322] 零钱兑换
#
from typing import *
# @lc code=start
# 关键是理解状态转移方程，最大需要回溯到 max(coins)
# 需要分清楚你是需要减 i 还是减 dp[i]
# 还有是需要考虑某些特殊的情况 [2] 4
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [0]
        coinsset = set(coins)
        for i in range(1,amount+1):
            if i in coinsset:
                dp.append(1)
            else:
                previous = [dp[i-j]for j in coins if i-j>=0 and  dp[i-j] > 0]
                if previous:
                    dp.append(min(previous)+1)
                else:
                    dp.append(-1)
        return  dp[-1] 
        ############ 这种回溯太慢了 ##################
                
                # for j in range(max(i-k,0),i):
                #     if i - j in coins:
                #         current = min(dp[j],current)
                # dp.append(current+1)
# @lc code=end



if __name__ == "__main__":
    print(Solution().coinChange([1,2,5],11))
    print(Solution().coinChange([2],4))
    print(Solution().coinChange([2],3))
    # print(Solution().coinChange([122,112,383,404,25,368],6977))