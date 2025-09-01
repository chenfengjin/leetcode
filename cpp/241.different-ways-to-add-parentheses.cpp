/*
 * @lc app=leetcode id=241 lang=cpp
 *
 * [241] Different Ways to Add Parentheses
 *
 * https://leetcode.com/problems/different-ways-to-add-parentheses/description/
 *
 * algorithms
 * Medium (72.59%)
 * Likes:    6239
 * Dislikes: 389
 * Total Accepted:    356K
 * Total Submissions: 490.4K
 * Testcase Example:  '"2-1-1"'
 *
 * Given a string expression of numbers and operators, return all possible
 * results from computing all the different possible ways to group numbers and
 * operators. You may return the answer in any order.
 * 
 * The test cases are generated such that the output values fit in a 32-bit
 * integer and the number of different results does not exceed 10^4.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: expression = "2-1-1"
 * Output: [0,2]
 * Explanation:
 * ((2-1)-1) = 0 
 * (2-(1-1)) = 2
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: expression = "2*3-4*5"
 * Output: [-34,-14,-10,-10,10]
 * Explanation:
 * (2*(3-(4*5))) = -34 
 * ((2*3)-(4*5)) = -14 
 * ((2*(3-4))*5) = -10 
 * (2*((3-4)*5)) = -10 
 * (((2*3)-4)*5) = 10
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= expression.length <= 20
 * expression consists of digits and the operator '+', '-', and '*'.
 * All the integer values in the input expression are in the range [0, 99].
 * The integer values in the input expression do not have a leading '-' or '+'
 * denoting the sign.
 * 
 * 
 */

#include "leetcode.h"
// @lc code=start
class Solution {
private:
    map<string, vector<int>> memory;

public:
    vector<int> diffWaysToCompute(string s) {
        if (memory.find(s)!=memory.end()){
            return memory[s];
        }
        bool is_number = true;
        for (char ch : s)
        {
            if (ch == '+' || ch == '-' || ch == '*' || ch == '/')
            {
                is_number = false;
                break;
            }
       
        }
        if (is_number){
          int value = std::stoi(s);
          return {value};
        }

        vector<int> result= vector<int>();
        for (int i = 0; i < s.length();i++){
            char ch = s[i];
            if (isdigit(ch))
            {
                continue;
            }
            vector<int> left = diffWaysToCompute(s.substr(0, i));
            vector<int> right = diffWaysToCompute(s.substr(i + 1));

            for (int i : left)
            {
                for(int j:right){
                  switch (ch) {
                  case '+': {
                      result.push_back(i + j);
                      break;
                  }
                  case '-': {
                    result.push_back(i - j);
                    break;
                  }
                  case '*': {
                    result.push_back(i * j);
                    break;
                  }
                  case '/':
                    result.push_back(i / j);
                    break;
                  }
                }
            }
        }
        memory[s] = result;
        return result;
    }
};
// @lc code=end

TEST(diffWaysToCompute,case1){
    
    print_vector(Solution().diffWaysToCompute("2-1-1"));
}
