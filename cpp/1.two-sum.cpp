/*
 * @lc app=leetcode id=1 lang=cpp
 *
 * [1] Two Sum
 */
#include <iostream>
#include <vector>
#include <map>
using namespace std;


// @lc code=start
class Solution {
public:
  vector<int> twoSum(vector<int> &nums, int target) {
    map<int, int> mem;
    for (int i = 0; i < nums.size();i++){
      auto num = nums[i];
      if (mem.count(num)!=0){
        return vector<int>{mem[num], i};
      }else
      {
        mem[target - num] = i;
      }
    }
  return std::vector<int>{0, 0};
  }

};
// @lc code=end
int main(){
  auto nums = std::vector<int>{3,2,4};
  auto target = 6;
  auto ret = Solution().twoSum(nums, target);
  cout  << ret[0] << "," << ret[1] << endl;
  return 0;
}
