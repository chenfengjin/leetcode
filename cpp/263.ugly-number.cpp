/*
 * @lc app=leetcode id=263 lang=cpp
 *
 * [263] Ugly Number
 *
 * https://leetcode.com/problems/ugly-number/description/
 *
 * algorithms
 * Easy (42.51%)
 * Likes:    3675
 * Dislikes: 1770
 * Total Accepted:    669.6K
 * Total Submissions: 1.6M
 * Testcase Example:  '6'
 *
 * An ugly number is a positive integer which does not have a prime factor
 * other than 2, 3, and 5.
 *
 * Given an integer n, return true if n is an ugly number.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 6
 * Output: true
 * Explanation: 6 = 2 × 3
 *
 *
 * Example 2:
 *
 *
 * Input: n = 1
 * Output: true
 * Explanation: 1 has no prime factors.
 *
 *
 * Example 3:
 *
 *
 * Input: n = 14
 * Output: false
 * Explanation: 14 is not ugly since it includes the prime factor 7.
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= n <= 2^31 - 1
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  bool isUgly(int n) {
    if (n <= 0) {
      return false;
    }

    while (n % 2 == 0) {
      n /= 2;
    }
    while (n % 3 == 0) {
      n /= 3;
    }
    while (n % 5 == 0) {
      n /= 5;
    }
    return n == 1;
  }
};
// @lc code=end
