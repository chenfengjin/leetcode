#
# @lc app=leetcode id=322 lang=python3
#
# [322] Coin Change
#

# @lc code=start
from typing import *
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [0]
        dp.extend([-1]* (amount))
        for i in range(1,amount + 1 ):
            min_value = -1
            for coin in coins:
                if not i -coin < 0 and not dp[i-coin] == -1 :
                    if min_value == -1:
                        min_value = dp[i-coin] + 1 
                        continue
                    if  dp[i-coin] + 1  <  min_value:
                        min_value = dp[i-coin] + 1 
            dp[i]=min_value
        print(dp)
        return dp[-1]

# @lc code=end
if __name__ == "__main__":
    print(Solution().coinChange([186,419,83,408],6249))

