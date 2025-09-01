/*
 * @lc app=leetcode id=228 lang=cpp
 *
 * [228] Summary Ranges
 *
 * https://leetcode.com/problems/summary-ranges/description/
 *
 * algorithms
 * Easy (53.30%)
 * Likes:    4375
 * Dislikes: 2340
 * Total Accepted:    833.8K
 * Total Submissions: 1.6M
 * Testcase Example:  '[0,1,2,4,5,7]'
 *
 * You are given a sorted unique integer array nums.
 * 
 * A range [a,b] is the set of all integers from a to b (inclusive).
 * 
 * Return the smallest sorted list of ranges that cover all the numbers in the
 * array exactly. That is, each element of nums is covered by exactly one of
 * the ranges, and there is no integer x such that x is in one of the ranges
 * but not in nums.
 * 
 * Each range [a,b] in the list should be output as:
 * 
 * 
 * "a->b" if a != b
 * "a" if a == b
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [0,1,2,4,5,7]
 * Output: ["0->2","4->5","7"]
 * Explanation: The ranges are:
 * [0,2] --> "0->2"
 * [4,5] --> "4->5"
 * [7,7] --> "7"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [0,2,3,4,6,8,9]
 * Output: ["0","2->4","6","8->9"]
 * Explanation: The ranges are:
 * [0,0] --> "0"
 * [2,4] --> "2->4"
 * [6,6] --> "6"
 * [8,9] --> "8->9"
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 0 <= nums.length <= 20
 * -2^31 <= nums[i] <= 2^31 - 1
 * All the values of nums are unique.
 * nums is sorted in ascending order.
 * 
 * 
 */

#include "leetcode.h"
// @lc code=start

class Solution {
public:
  vector<string> summaryRanges(vector<int> &nums) {
    vector<string> ret;
    int i = 0;
    int n = nums.size();
    while (i < n) {
      int low = i;
      i++;
      while (i < n && nums[i] == nums[i - 1] + 1) {
        i++;
      }
      int high = i - 1;
      string temp = to_string(nums[low]);
      if (low < high) {
        temp.append("->");
        temp.append(to_string(nums[high]));
      }
      ret.push_back(move(temp));
    }
    return ret;
  }
};

          
class
      Solution1 {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        if (nums.size()==0){
            return vector<string>();
        }
        if (nums.size()==1){
            return vector<string>{to_string(nums[0])};
        }
        
        //  把第一个元素也放到 nums 里，是可以不用处理最后一段
        //  由于是升序且我们处理了只有一个元素的情况，一定不等于 nums 最后一个元素加 1

        nums.push_back(nums[0]);
        vector<string> ret = vector<string>();

        int start = nums[0];
        for (int i = 1; i < nums.size();i++){

          if (nums[i] ==  nums[i - 1] +1){ 
                continue;
            }
            // 这里很容易漏掉
            if (nums[i - 1] - start == 0) {
              ret.push_back(to_string(start));
            } else {
              ret.push_back(to_string(start) + "->" + to_string(nums[i - 1]));
            }
            start = nums[i];
        }
        return ret;
    }
};
// @lc code=end


// 