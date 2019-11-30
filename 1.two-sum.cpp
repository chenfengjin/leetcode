/*
 * @lc app=leetcode id=1 lang=cpp
 *
 * [1] Two Sum
 */
# include<vector>
# include<map>
# include<set>
using namespace std;
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        set<int> s;
        vector<int> target_result;
        for(vector<int>::iterator i = nums.begin(); i != nums.end();i++){
            s.insert(target - *i);
            set<int>::iterator b;
            b=s.find(*i);
            if(b!=s.end());
                target_result.push_back(*i);
                target_result.push_back(target - *i);
                return target_result;
            }
        return target_result;

        }
};

