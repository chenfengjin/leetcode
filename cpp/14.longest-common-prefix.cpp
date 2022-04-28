/*
 * @lc app=leetcode id=14 lang=cpp
 *
 * [14] Longest Common Prefix
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        int minLength = 200;
        for (auto str:strs){
            if(str.size()<minLength){
                minLength = str.size();
            }
        }
        for (int i = 0; i < minLength;i++){
            char ch = strs[0][i];
            for (auto str : strs)
            {
                if (str[i]!=ch){
                    return str.substr(0, i);
                }
            }
        }
        return strs[0].substr(0, minLength);
    }
};
// @lc code=end

int main(){
    vector<string> s{"a", "ab"};
    cout << Solution().longestCommonPrefix(s);
}