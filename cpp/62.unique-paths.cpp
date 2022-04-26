/*
 * @lc app=leetcode id=62 lang=cpp
 *
 * [62] Unique Paths
 */
#include<leetcode-helper.h>

// @lc code=start
class Solution {
public:
    int uniquePaths(int m, int n) {
        int dp[m][n];
        for (auto i = 0; i < m;i++){
            for (auto j = 0; j < n;j++){
                if (i==0||j==0){
                    dp[i][j] = 1;
                }else{
                    dp[i][j] = 0;
                }
            }
        }
        for (auto i = 1; i < m;i++){
            for (auto j = 1; j < n;j++){
                dp[i][j] = dp[i][j - 1] + dp[i - 1][j];
            }
        }
        return dp[m - 1][n - 1];
    }
};
// @lc code=end

