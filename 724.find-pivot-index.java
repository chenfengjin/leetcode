/*
 * @lc app=leetcode id=724 lang=java
 *
 * [724] Find Pivot Index
 */
// @lc code=start
class Solution {
    public int pivotIndex(int[] nums) {
        int sum = 0;
        int subsum = 0;
        for(int i = 0;i < nums.length;i++){
            sum += nums[i];
        }
        for(int i = 0;i  < nums.length;i++){
            if (subsum+subsum+nums[i] == sum){
                return  i;
            }
            subsum += nums[i];
        }
        return -1;
    }
}
// @lc code=end

