/*
 * @lc app=leetcode id=139 lang=cpp
 *
 * [139] Word Break
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    bool wordBreak(string s, vector<string>& wordDict) {
        // default??
        bool dp[s.size()];
        dp[0] = true;
        for (int i = 0;i<s.size();i++){
            for (auto word:wordDict){
                // TODO
                if (i-word.size()>0&&dp[i-word.size()]&&s.substr(i-word.size(),i)==word){
                    dp[i] = true;
                    continue;
                }
            }
        }
        return dp[s.size() - 1];
    }
};
// @lc code=end

