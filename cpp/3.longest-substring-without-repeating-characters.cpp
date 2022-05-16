/*
 * @lc app=leetcode id=3 lang=cpp
 *
 * [3] Longest Substring Without Repeating Characters
 */
#include<string>
#include<map>
#include<iostream>
using namespace std;
// @lc code=start
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        map<char, int> charCounts;
        auto left = 0;
        auto right = 0;
        auto max = 0;
        while (right < s.size())
        {
            // printMap(charCounts);
            if (charCounts.count(s.at(right)))
            {
                if (max < right-left){
                    max = right - left;
                }
                left++;
                
                // small trick
                while (left == 0 || s.at(left - 1) != s.at(right))
                {
                    charCounts[s.at(left)] = 0;
                    left++;
                }
            }
            else
            {
                charCounts[s.at(right)] = 1;
            }
            right++;
        }
          if (max < right-left){
            max = right - left;
        }
        return max;
    }
private:
    void printMap(map<char,int>m){
        for (auto iter = m.begin(); iter != m.end();++iter){
            cout << iter->first << " " << iter->second << endl;
        }
    }
};
// @lc code=end

int main(){
    cout << Solution().lengthOfLongestSubstring(" ");
}