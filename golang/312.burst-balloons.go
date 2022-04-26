/*
 * @lc app=leetcode id=312 lang=golang
 *
 * [312] Burst Balloons
 */

// @lc code=start
func maxCoins(nums []int) int {
    
}
// @lc code=end

class Solution {
public:
    int maxCoins(vector<int>& nums) {
        int size = nums.size();
        vector<int> scores(size + 2, 1);
        for (int idx = 0; idx < size; ++idx) {
            scores[idx + 1] = nums[idx];
        }
        vector<vector<int>> dp(size + 2, vector<int>(size + 2, 0));
        for (int k = 1; k <= size; ++k) {
            for (int i = 1; i + k - 1 <= size; ++i) {
                for (int j = i; j <= i + k - 1; ++j) {
                    dp[i][i + k - 1] = max(dp[i][i + k - 1], dp[i][j - 1] + dp[j + 1][i + k - 1] + scores[i - 1] * scores[j] * scores[i + k]);
                }
            }
        }
        return dp[1][size];
    }
};

