/*
 * @lc app=leetcode id=121 lang=cpp
 *
 * [121] Best Time to Buy and Sell Stock
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        auto min = 1000000;
        auto max = 0;
        for (auto price : prices)
        {
            if (min>price){
                min = price;
            }
            if (price-min>max){
                max = price - min;
            }
        }
        return max;
    }
};
// @lc code=end

