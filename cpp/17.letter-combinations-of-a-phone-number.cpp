/*
 * @lc app=leetcode id=17 lang=cpp
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (64.38%)
 * Likes:    20171
 * Dislikes: 1101
 * Total Accepted:    2.7M
 * Total Submissions: 4.2M
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent. Return the answer in
 * any order.
 *
 * A mapping of digits to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 * Example 1:
 *
 *
 * Input: digits = "23"
 * Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
 *
 *
 * Example 2:
 *
 *
 * Input: digits = ""
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: digits = "2"
 * Output: ["a","b","c"]
 *
 *
 *
 * Constraints:
 *
 *
 * 0 <= digits.length <= 4
 * digits[i] is a digit in the range ['2', '9'].
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  vector<string> letterCombinations(string digits) {
    if (!digits.size()) {
      return vector<string>();
    }
    auto cur = "";
    int start = 0;
    int length = digits.length();
    auto table = map<char, string>();
    {
      table.insert(pair<char, string>('2', "abc"));
      table.insert(pair<char, string>('3', "def"));
      table.insert(pair<char, string>('4', "ghi"));
      table.insert(pair<char, string>('5', "jkl"));
      table.insert(pair<char, string>('6', "mno"));
      table.insert(pair<char, string>('7', "pqrs"));
      table.insert(pair<char, string>('8', "tuv"));
      table.insert(pair<char, string>('9', "wxyz"));
    }
    auto collector = vector<string>();
    helper(digits, cur, start, length, table, collector);
    return collector;
  }

  void helper(string digits, string cur, int start, int length,
              map<char, string> table, vector<string> &collector) {
    if (start == length) {
      collector.push_back(cur);
      return;
    }
    for (char ch : table[digits[start]]) {
      cur.push_back(ch);
      helper(digits, cur, start + 1, length, table, collector);
      cur.pop_back();
    }
  }
};
// @lc code=end
