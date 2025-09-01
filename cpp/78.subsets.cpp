/*
 * @lc app=leetcode id=78 lang=cpp
 *
 * [78] Subsets
 *
 * https://leetcode.com/problems/subsets/description/
 *
 * algorithms
 * Medium (81.32%)
 * Likes:    18573
 * Dislikes: 323
 * Total Accepted:    2.7M
 * Total Submissions: 3.3M
 * Testcase Example:  '[1,2,3]'
 *
 * Given an integer array nums of unique elements, return all possible subsets
 * (the power set).
 *
 * The solution set must not contain duplicate subsets. Return the solution in
 * any order.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3]
 * Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0]
 * Output: [[],[0]]
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10
 * -10 <= nums[i] <= 10
 * All the numbers of nums are unique.
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  vector<vector<int>> subsets(vector<int> &nums) {
    vector<vector<int>> collector = vector<vector<int>>();
    for (int i = 0; i < 1 << nums.size(); i++) {
      vector<int> cur = vector<int>();
      int n = i;
      int idx = 0;
      while (n) {
        if (n & 1) {
          cur.push_back(nums[idx]);
        }
        idx++;
        n >>= 1;
      }
      collector.push_back(cur);
    }
    return collector;
  }
};
// @lc code=end

TEST(Subset, case) {
  vector<vector<int>> want = {{},  {1},    {2},    {1, 2},
                              {3}, {1, 3}, {2, 3}, {1, 2, 3}};
  vector<int> testCase = {1, 2, 3};
  auto got = Solution().subsets(testCase);
  ASSERT_EQ(got, want);
}
