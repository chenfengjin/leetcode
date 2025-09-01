/*
 * @lc app=leetcode id=53 lang=cpp
 *
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Medium (52.38%)
 * Likes:    36420
 * Dislikes: 1543
 * Total Accepted:    5.3M
 * Total Submissions: 10.2M
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the subarray with the largest sum, and
 * return its sum.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: The subarray [4,-1,2,1] has the largest sum 6.
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1]
 * Output: 1
 * Explanation: The subarray [1] has the largest sum 1.
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [5,4,-1,7,8]
 * Output: 23
 * Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 *
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution using the divide and conquer approach, which is more subtle.
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  int maxSubArray(vector<int> &nums) {
    auto dp = vector<int>(nums.size());
    dp[0] = nums[0];
    int ret = dp[0];

    for (int i = 1; i < nums.size(); i++) {
      dp[i] = dp[i - 1] < 0 ? nums[i] : dp[i - 1] + nums[i];
      if (dp[i] > ret) {
        ret = dp[i];
      }
    }
    return ret;
  }
};
// @lc code=end
