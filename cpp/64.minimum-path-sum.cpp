/*
 * @lc app=leetcode id=64 lang=cpp
 *
 * [64] Minimum Path Sum
 *
 * https://leetcode.com/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (66.99%)
 * Likes:    13298
 * Dislikes: 186
 * Total Accepted:    1.6M
 * Total Submissions: 2.4M
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * Given a m x n grid filled with non-negative numbers, find a path from top
 * left to bottom right, which minimizes the sum of all numbers along its
 * path.
 *
 * Note: You can only move either down or right at any point in time.
 *
 *
 * Example 1:
 *
 *
 * Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
 * Output: 7
 * Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
 *
 *
 * Example 2:
 *
 *
 * Input: grid = [[1,2,3],[4,5,6]]
 * Output: 12
 *
 *
 *
 * Constraints:
 *
 *
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  int minPathSum(vector<vector<int>> &grid) {
    // m 行数
    // n 列数
    int m = grid.size();
    int n = grid[0].size();
    auto dp = vector<vector<int>>(m);
    for (int i = 0; i < m; i++) {
      dp[i] = vector<int>(n);
    }
    dp[0][0] = grid[0][0];
    for (int i = 1; i < m; i++) {
      dp[i][0] = dp[i - 1][0] + grid[i][0];
    }
    for (int j = 1; j < n; j++) {
      dp[0][j] = dp[0][j - 1] + grid[0][j];
    }
    for (int i = 1; i < m; i++) {
      for (int j = 1; j < n; j++) {
        dp[i][j] = min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j];
      }
    }
    return dp[m - 1][n - 1];
  }
};
// @lc code=end

TEST(minPathSum, case1) {
  auto a = vector<vector<int>>(3);
  a[0] = vector<int>({1, 3, 1});
  a[1] = vector<int>({1, 5, 1});
  a[2] = vector<int>({4, 2, 1});
  Solution().minPathSum(a);
}
// 1 3 1
// 1 5 1
// 4 2 1
