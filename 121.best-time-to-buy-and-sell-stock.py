#
# @lc app=leetcode id=121 lang=python3
#
# [121] Best Time to Buy and Sell Stock
#
from typing import *
# @lc code=start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        #  it is ok but too slow
        # return 0 if len(prices)<2 else max(max([prices[i] - min(prices[:i]) for i in range(1,len(prices))]),0)

        if len(prices)< 2:
            return 0
        min_array = [0]*len(prices)
        min_array[0] = prices[0]
        max_profit = 0
        for i in range(1,len(prices)):
            min_array[i] = min(min_array[i-1],prices[i])
            max_profit = max(max_profit,prices[i]-min_array[i])
        return max_profit

        
# @lc code=end


if __name__ == "__main__":
    print(Solution().maxProfit([7,1,5,3,6,4]))