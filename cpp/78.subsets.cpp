/*
 * @lc app=leetcode id=78 lang=cpp
 *
 * [78] Subsets
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<vector<int>> ret;
        auto size = nums.size();
        for (auto i = 0; i < 1 << size;i++){
            auto tmp = i;
            auto cur = 0;
            vector<int> item;
            while (tmp > 0)
            {
                if (tmp&1){
                    item.push_back(nums[cur]);
                }
                cur++;
                tmp >>= 1;
            }
            ret.push_back(item);
        }
        return ret;
    }
};
// @lc code=end

int main(){
    vector<int> nums{1, 2, 3};
    auto result = Solution().subsets(nums);
    for (auto iter:result)
    {
        for(auto i:iter){
            cout << i;
        }
        cout << endl;

    }
}