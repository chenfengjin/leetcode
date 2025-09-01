/*
 * @lc app=leetcode id=14 lang=cpp
 *
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (46.02%)
 * Likes:    19852
 * Dislikes: 4798
 * Total Accepted:    4.9M
 * Total Submissions: 10.6M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 *
 * Example 1:
 *
 *
 * Input: strs = ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: strs = ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] consists of only lowercase English letters if it is non-empty.
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  string longestCommonPrefix(vector<string> &strs) {
    if (strs.size() == 0) {
      return "";
    }
    if (strs.size() == 1) {
      return strs[0];
    }
    int minLength = 200;
    for (auto str : strs) {
      if (minLength > str.size()) {
        minLength = str.size();
      }
    }
    int i = 0;
    for (; i < minLength; i++) {
      for (int j = 1; j < strs.size(); j++) {
        if (strs[j][i] != strs[0][i]) {
          return strs[0].substr(0, i);
        }
      }
    }
    return strs[0].substr(0, minLength);
  }
};
// @lc code=end

TEST(longestCommonPrefix, case1) {
  auto testCase = vector<string>{"ab", "a"};
  Solution().longestCommonPrefix(testCase);
}
