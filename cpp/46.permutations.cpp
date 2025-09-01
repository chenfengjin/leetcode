/*
 * @lc app=leetcode id=46 lang=cpp
 *
 * [46] Permutations
 *
 * https://leetcode.com/problems/permutations/description/
 *
 * algorithms
 * Medium (81.01%)
 * Likes:    20272
 * Dislikes: 366
 * Total Accepted:    2.7M
 * Total Submissions: 3.4M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an array nums of distinct integers, return all the possible
 * permutations. You can return the answer in any order.
 *
 *
 * Example 1:
 * Input: nums = [1,2,3]
 * Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * Example 2:
 * Input: nums = [0,1]
 * Output: [[0,1],[1,0]]
 * Example 3:
 * Input: nums = [1]
 * Output: [[1]]
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 6
 * -10 <= nums[i] <= 10
 * All the integers of nums are unique.
 *
 *
 */

#include "leetcode.h"
// @lc code=start
//  key point of using traceback to solve permutation
// 1. every recursion start from zero
// 2. have an memory to remove duplication
class Solution {
public:
  vector<vector<int>> permute(vector<int> &nums) {
    vector<vector<int>> collector = vector<vector<int>>();
    vector<int> cur = vector<int>();
    int start = 0;
    auto seen = vector<int>(nums.size(), 0);
    helper(nums, cur, start, collector, seen);

    return collector;
  }
  void helper(vector<int> &nums, vector<int> &cur, int start,
              vector<vector<int>> &collector, vector<int> &seen) {
    if (cur.size() == nums.size()) {
      collector.push_back(cur);
      return;
    }

    for (int i = 0; i < nums.size(); i++) {
      if (seen[i]) {
        continue;
      }
      cur.push_back(nums[i]);
      seen[i] = 1;
      helper(nums, cur, i + 1, collector, seen);
      seen[i] = 0;
      cur.pop_back();
    }
  }
};
// @lc code=end

TEST(permute, case1) {
  vector<int> testCase = {1, 2, 3};
  auto got = Solution().permute(testCase);
  // cout << "got size:" << got.size() << endl;
  vector<vector<int>> want = {{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
                              {2, 3, 1}, {3, 1, 2}, {3, 2, 1}};
  ASSERT_EQ(got, want);
}