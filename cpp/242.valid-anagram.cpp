/*
 * @lc app=leetcode id=242 lang=cpp
 *
 * [242] Valid Anagram
 *
 * https://leetcode.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (67.06%)
 * Likes:    13444
 * Dislikes: 451
 * Total Accepted:    5.2M
 * Total Submissions: 7.7M
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * Given two strings s and t, return true if t is an anagram of s, and false
 * otherwise.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "anagram", t = "nagaram"
 * 
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "rat", t = "car"
 * 
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length, t.length <= 5 * 10^4
 * s and t consist of lowercase English letters.
 * 
 * 
 * 
 * Follow up: What if the inputs contain Unicode characters? How would you
 * adapt your solution to such a case?
 * 
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.length()!=t.length()){
            return false;
        }
        map<char, int> m = map<char, int>();
        for (char c : s)
        {
            m[c]++;
        }
        for(char c:t){
            m[c]--;
            if (m[c]<0){
                return false;
            }
        }
        for (auto iter = m.begin(); iter != m.end();iter++){
            if (iter->second!=0){
                return false;
            }
        }
        return true;
    };
};
// @lc code=end
