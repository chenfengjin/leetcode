/*
 * @lc app=leetcode id=9 lang=cpp
 *
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (59.55%)
 * Likes:    14605
 * Dislikes: 2850
 * Total Accepted:    6.9M
 * Total Submissions: 11.6M
 * Testcase Example:  '121'
 *
 * Given an integer x, return true if x is a palindrome, and false
 * otherwise.
 *
 *
 * Example 1:
 *
 *
 * Input: x = 121
 * Output: true
 * Explanation: 121 reads as 121 from left to right and from right to left.
 *
 *
 * Example 2:
 *
 *
 * Input: x = -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: x = 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 *
 * Follow up: Could you solve it without converting the integer to a string?
 */
//  TODO more effective
#include "leetcode.h"
// @lc code=start
class Solution {
public:
  bool isPalindrome(int x) {
    if (x < 0) {
      return false;
    }
    auto digits = vector<int>();
    while (x) {
      digits.push_back(x % 10);
      x /= 10;
    }
    for (int i = 0; i < digits.size(); i++) {
      if (digits[i] != digits[digits.size() - i - 1]) {
        return false;
      }
    }
    return true;
  }
};
// @lc code=end
