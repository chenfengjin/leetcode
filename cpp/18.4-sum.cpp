/*
 * @lc app=leetcode id=18 lang=cpp
 *
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum/description/
 *
 * algorithms
 * Medium (38.76%)
 * Likes:    12335
 * Dislikes: 1481
 * Total Accepted:    1.4M
 * Total Submissions: 3.6M
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array nums of n integers, return an array of all the unique
 * quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
 *
 *
 * 0 <= a, b, c, d < n
 * a, b, c, and d are distinct.
 * nums[a] + nums[b] + nums[c] + nums[d] == target
 *
 *
 * You may return the answer in any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,0,-1,0,-2,2], target = 0
 * Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,2,2,2,2], target = 8
 * Output: [[2,2,2,2]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 200
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 *
 *
 */
//  TODO more effective
#include "leetcode.h"
// @lc code=start
class Solution {
public:
  vector<vector<int>> fourSum(vector<int> &nums, int target) {
    return nSum(nums, 4, target);
  }
  vector<vector<int>> nSum(vector<int> &nums, int n, int target) {
    std::sort(std ::begin(nums), std::end(nums));
    int length = nums.size();
    auto collector = vector<vector<int>>();
    if (length < n) {
      return collector;
    }
    auto seen = vector<int>(length, 0);
    int start = 0;
    auto cur = vector<int>();
    helper(nums, seen, cur, start, length, collector, n, target);
    return collector;
  }

  void helper(vector<int> &nums, vector<int> &seen, vector<int> &cur, int start,
              int length, vector<vector<int>> &collector, int n, int target) {

    if (cur.size() == n) {
      long long total = 0;
      for (auto num : cur) {
        total += num;
      }
      if (total == target) {
        collector.push_back(cur);
      }
      return;
    }

    for (int i = start; i < length; i++) {
      if (i > 0 && nums[i] == nums[i - 1] && !seen[i - 1]) {
        continue;
      }
      seen[i] = 1;
      cur.push_back(nums[i]);
      helper(nums, seen, cur, i + 1, length, collector, n, target);

      seen[i] = 0;
      cur.pop_back();
    }
  }
};
// @lc code=end

TEST(FourSum, case1) {
  vector<int> testCase = {1000000000, 1000000000, 1000000000, 1000000000};
  Solution().fourSum(testCase, 0);
}