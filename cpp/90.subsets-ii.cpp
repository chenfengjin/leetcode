/*
 * @lc app=leetcode id=90 lang=cpp
 *
 * [90] Subsets II
 *
 * https://leetcode.com/problems/subsets-ii/description/
 *
 * algorithms
 * Medium (60.01%)
 * Likes:    10553
 * Dislikes: 386
 * Total Accepted:    1.3M
 * Total Submissions: 2.2M
 * Testcase Example:  '[1,2,2]'
 *
 * Given an integer array nums that may contain duplicates, return all possible
 * subsets (the power set).
 *
 * The solution set must not contain duplicate subsets. Return the solution in
 * any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,2]
 * Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
 * Example 2:
 * Input: nums = [0]
 * Output: [[],[0]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10
 * -10 <= nums[i] <= 10
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  vector<vector<int>> subsetsWithDup(vector<int> &nums) {
    if (nums.size() == 0) {
      return vector<vector<int>>();
    }
    auto ret = vector<vector<int>>();
    std::sort(std ::begin(nums), std::end(nums));
    int start = 0;
    auto seen = vector<int>(nums.size(), 0);
    auto cur = vector<int>();
    helper(nums, start, seen, cur, ret);

    return ret;
  }
  void helper(vector<int> &nums, int start, vector<int> &seen, vector<int> &cur,
              vector<vector<int>> &collect) {
    collect.push_back(cur);

    for (int i = start; i < nums.size(); i++) {
      if (i > 0 && nums[i] == nums[i - 1] && !seen[i - 1]) {
        continue;
      }
      seen[i] = 1;
      cur.push_back(nums[i]);
      helper(nums, i + 1, seen, cur, collect);

      cur.pop_back();

      seen[i] = 0;
    }
  }
};
// @lc code=end

TEST(subset2, case1) {
  auto testCase = vector<int>({1, 2, 2});
  auto results = Solution().subsetsWithDup(testCase);
  for (auto result : results) {
    for (auto num : result) {
      cout << num << "->";
    }
    cout << endl;
  }
}