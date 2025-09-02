/*
 * @lc app=leetcode id=216 lang=cpp
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (72.29%)
 * Likes:    6463
 * Dislikes: 120
 * Total Accepted:    710K
 * Total Submissions: 982K
 * Testcase Example:  '3\n7'
 *
 * Find all valid combinations of k numbers that sum up to n such that the
 * following conditions are true:
 *
 *
 * Only numbers 1 through 9 are used.
 * Each number is used at most once.
 *
 *
 * Return a list of all possible valid combinations. The list must not contain
 * the same combination twice, and the combinations may be returned in any
 * order.
 *
 *
 * Example 1:
 *
 *
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 * Explanation:
 * 1 + 2 + 4 = 7
 * There are no other valid combinations.
 *
 * Example 2:
 *
 *
 * Input: k = 3, n = 9
 * Output: [[1,2,6],[1,3,5],[2,3,4]]
 * Explanation:
 * 1 + 2 + 6 = 9
 * 1 + 3 + 5 = 9
 * 2 + 3 + 4 = 9
 * There are no other valid combinations.
 *
 *
 * Example 3:
 *
 *
 * Input: k = 4, n = 1
 * Output: []
 * Explanation: There are no valid combinations.
 * Using 4 different numbers in the range [1,9], the smallest sum we can get is
 * 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= k <= 9
 * 1 <= n <= 60
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  vector<vector<int>> combinationSum3(int k, int n) {
    auto nums = vector<int>({1, 2, 3, 4, 5, 6, 7, 8, 9});
    return combinationSum2(nums, n, k);
  }

  vector<vector<int>> combinationSum2(vector<int> &candidates, int target,
                                      int k) {
    auto cur = vector<int>();
    auto ret = vector<vector<int>>();
    int start = 0;
    auto seen = vector<int>(candidates.size());
    // key point 1: sort for pruning
    sort(candidates.begin(), candidates.end());
    helper(candidates, cur, target, ret, start, seen, k);
    return ret;
  }

  void helper(vector<int> &nums, vector<int> cur, int target,
              vector<vector<int>> &collector, int start, vector<int> &seen,
              int k) {
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
        if (num == target && cur.size() == k - 1) {
          cur.push_back(num);
          collector.push_back(cur);
          cur.pop_back();
          return;
        }
        cur.push_back(num);
        seen[i] = 1;
        // key point3: start from i+1 to disallow reuse of elements
        helper(nums, cur, target - num, collector, i + 1, seen, k);
        seen[i] = 0;
        cur.pop_back();
      }
    }
  }
};
// @lc code=end
