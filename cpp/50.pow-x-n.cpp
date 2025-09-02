/*
 * @lc app=leetcode id=50 lang=cpp
 *
 * [50] Pow(x, n)
 *
 * https://leetcode.com/problems/powx-n/description/
 *
 * algorithms
 * Medium (37.44%)
 * Likes:    11045
 * Dislikes: 10322
 * Total Accepted:    2.4M
 * Total Submissions: 6.3M
 * Testcase Example:  '2.00000\n10'
 *
 * Implement pow(x, n), which calculates x raised to the power n (i.e.,
 * x^n).
 *
 *
 * Example 1:
 *
 *
 * Input: x = 2.00000, n = 10
 * Output: 1024.00000
 *
 *
 * Example 2:
 *
 *
 * Input: x = 2.10000, n = 3
 * Output: 9.26100
 *
 *
 * Example 3:
 *
 *
 * Input: x = 2.00000, n = -2
 * Output: 0.25000
 * Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25
 *
 *
 *
 * Constraints:
 *
 *
 * -100.0 < x < 100.0
 * -2^31 <= n <= 2^31-1
 * n is an integer.
 * Either x is not zero or n > 0.
 * -10^4 <= x^n <= 10^4
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  double myPow(double x, int n) {
    // deal with overflow problem
    if (-2147483648 == n) {
      return x > 1 ? 0 : 1;
    }
    int retSign = x < 0 && n & 1;
    int nSign = n < 0;
    x = x < 0 ? -1 * x : x;
    n = n < 0 ? -1 * n : n;

    double ret = 1.0;
    while (n) {
      if (n & 1) {
        ret *= x;
      }
      x = x * x;
      n >>= 1;
    }
    ret = nSign ? 1 / ret : ret;
    ret = retSign ? -1 * ret : ret;
    return ret;
  }
};
// @lc code=end
