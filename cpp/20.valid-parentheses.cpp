/*
 * @lc app=leetcode id=20 lang=cpp
 *
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (42.75%)
 * Likes:    26453
 * Dislikes: 1926
 * Total Accepted:    6.5M
 * Total Submissions: 15.2M
 * Testcase Example:  '"()"'
 *
 * Given a string s containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 * 
 * An input string is valid if:
 * 
 * 
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Every close bracket has a corresponding open bracket of the same type.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: s = "()"
 * 
 * Output: true
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: s = "()[]{}"
 * 
 * Output: true
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: s = "(]"
 * 
 * Output: false
 * 
 * 
 * Example 4:
 * 
 * 
 * Input: s = "([])"
 * 
 * Output: true
 * 
 * 
 * Example 5:
 * 
 * 
 * Input: s = "([)]"
 * 
 * Output: false
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= s.length <= 10^4
 * s consists of parentheses only '()[]{}'.
 * 
 * 
 */

#include "leetcode.h"
// @lc code=start
#include <stack>
using namespace std;
class Solution
{
public:
    bool isValid(string str) {
        stack<char> s= stack <char>();
            for (int i = 0; i < str.length(); i++)
        {
            char ch = str[i];
            switch(ch){
                case '(':
                case '[':
                case '{': {
                  s.push(ch);
                  break;
                }
                case ')':{
                    if(s.empty()||s.top()!='('){
                        return false;
                    }
                    s.pop();
                    break;
                }
                    break;
                case ']':
                {
                  if (s.empty()||s.top() != '[') {
                    return false;
                  }
                  s.pop();
                  break;
                }
                case '}': {
                  if (s.empty()|| s.top() != '{') {
                    return false;
                  }
                  s.pop();
                } break;
                }
        }
        return s.empty();
    }
};
// @lc code=end

TEST(isValid, case1) {
    ASSERT_TRUE(Solution().isValid("()"));
    ASSERT_TRUE(Solution().isValid("([])"));
    ASSERT_FALSE(Solution().isValid("([)]"));
}