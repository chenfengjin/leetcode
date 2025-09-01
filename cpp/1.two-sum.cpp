/*
 * @lc app=leetcode id=1 lang=cpp
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (56.20%)
 * Likes:    63871
 * Dislikes: 2329
 * Total Accepted:    18.6M
 * Total Submissions: 33M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers nums and an integer target, return indices of the
 * two numbers such that they add up to target.
 * 
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 * 
 * You can return the answer in any order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [2,7,11,15], target = 9
 * Output: [0,1]
 * Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [3,2,4], target = 6
 * Output: [1,2]
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: nums = [3,3], target = 6
 * Output: [0,1]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * Only one valid answer exists.
 * 
 * 
 * 
 * Follow-up: Can you come up with an algorithm that is less than O(n^2) time
 * complexity?
 */

 #include "leetcode.h"
// @lc code=start
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        map<int,int> diff = map<int, int>();
        vector<int> ret;
        for (int i = 0; i < nums.size(); i++)
        {
            auto num = nums[i];
            if (diff.find(num) != diff.end())
            {
                ret =  vector({diff.at(num),i});
                break;
            }
            diff[target - num] = i;
        }
        return ret;
    }
};
// @lc code=end



// 测试用例
TEST(ExampleTest, Case1) {
    vector<int> nums = {2, 7, 11, 15};
    int target = 9;
    vector<int> want = {0, 1};

    EXPECT_EQ(Solution().twoSum(nums,target),want);
}
