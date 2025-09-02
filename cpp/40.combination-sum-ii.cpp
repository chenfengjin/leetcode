/*
 * @lc app=leetcode id=40 lang=cpp
 *
 * [40] Combination Sum II
 *
 * https://leetcode.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (58.07%)
 * Likes:    11815
 * Dislikes: 363
 * Total Accepted:    1.5M
 * Total Submissions: 2.6M
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * Given a collection of candidate numbers (candidates) and a target number
 * (target), find all unique combinations in candidates where the candidate
 * numbers sum to target.
 *
 * Each number in candidates may only be used once in the combination.
 *
 * Note: The solution set must not contain duplicate combinations.
 *
 *
 * Example 1:
 *
 *
 * Input: candidates = [10,1,2,7,6,1,5], target = 8
 * Output:
 * [
 * [1,1,6],
 * [1,2,5],
 * [1,7],
 * [2,6]
 * ]
 *
 *
 * Example 2:
 *
 *
 * Input: candidates = [2,5,2,1,2], target = 5
 * Output:
 * [
 * [1,2,2],
 * [5]
 * ]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= candidates.length <= 100
 * 1 <= candidates[i] <= 50
 * 1 <= target <= 30
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  vector<vector<int>> combinationSum2(vector<int> &candidates, int target) {
    auto cur = vector<int>();
    auto ret = vector<vector<int>>();
    int start = 0;
    auto seen = vector<int>(candidates.size());
    // key point 1: sort for pruning
    sort(candidates.begin(), candidates.end());
    helper(candidates, cur, target, ret, start, seen);
    return ret;
  }

  void helper(vector<int> &nums, vector<int> cur, int target,
              vector<vector<int>> &collector, int start, vector<int> &seen) {
    // key point 2 start from i to get rid of duplication result
    for (int i = start; i < nums.size(); i++) {
      int num = nums[i];
      if (num > target) {
        return;
      } else {
        if (i > 0 && nums[i] == nums[i - 1] && !seen[i - 1]) {
          //  remove duplications in result
          continue;
        }
        if (num == target) {
          cur.push_back(num);
          collector.push_back(cur);
          cur.pop_back();
          return;
        }
        cur.push_back(num);
        seen[i] = 1;
        // key point3: start from i+1 to disallow reuse of elements
        helper(nums, cur, target - num, collector, i + 1, seen);
        seen[i] = 0;
        cur.pop_back();
      }
    }
  }
};
// @lc code=end
TEST(combinationSum2, case1) {
  auto testCase = vector<int>({2, 5, 2, 1, 2});
  Solution().combinationSum2(testCase, 5);
}