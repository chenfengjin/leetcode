/*
 * @lc app=leetcode id=16 lang=cpp
 *
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (47.18%)
 * Likes:    11174
 * Dislikes: 607
 * Total Accepted:    1.6M
 * Total Submissions: 3.3M
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * Given an integer array nums of length n and an integer target, find three
 * integers in nums such that the sum is closest to target.
 *
 * Return the sum of the three integers.
 *
 * You may assume that each input would have exactly one solution.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [-1,2,1,-4], target = 1
 * Output: 2
 * Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [0,0,0], target = 1
 * Output: 0
 * Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 =
 * 0).
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= nums.length <= 500
 * -1000 <= nums[i] <= 1000
 * -10^4 <= target <= 10^4
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  int threeSumClosest(vector<int> &nums, int target) {
    sort(nums.begin(), nums.end());
    int minDiff = 9999000;
    int cloestTarget;
    for (int i = 0; i < nums.size() - 2; i++) {
      //   int remain = target - nums[i];

      int l = i + 1;
      int r = nums.size() - 1;

      while (l < r) {
        int sum = nums[l] + nums[r] + nums[i];

        int diff = sum - target;

        diff = diff > 0 ? diff : -1 * diff;

        if (diff < minDiff) {
          minDiff = diff;
          cloestTarget = sum;
        }
        if (sum == target) {
          return target;
        } else if (sum > target) {
          r--;
        } else {
          l++;
        }
      }
    }
    return cloestTarget;
  }
};
// @lc code=end
TEST(ThreeSumColsest, case1) {
  vector<int> nums = {-1, 2, 1, -4};
  int target = 1;
  Solution().threeSumClosest(nums, target);
}