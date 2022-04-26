/*
 * @lc app=leetcode id=136 lang=cpp
 *
 * [136] Single Number
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    int singleNumber(vector<int>& nums) {
        int ret=0;
        for (auto iter = nums.begin();iter!=nums.end();iter++){
            ret ^= *iter;
        }
        return ret;
    }
};
// @lc code=end
int main(){
    vector<int> nums{2, 2, 1};
    cout << Solution().singleNumber(nums) <<endl;
}
