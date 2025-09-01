/*
 * @lc app=leetcode id=215 lang=cpp
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (68.31%)
 * Likes:    18176
 * Dislikes: 952
 * Total Accepted:    3.2M
 * Total Submissions: 4.7M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Given an integer array nums and an integer k, return the k^th largest
 * element in the array.
 *
 * Note that it is the k^th largest element in the sorted order, not the k^th
 * distinct element.
 *
 * Can you solve it without sorting?
 *
 *
 * Example 1:
 * Input: nums = [3,2,1,5,6,4], k = 2
 * Output: 5
 * Example 2:
 * Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
 * Output: 4
 *
 *
 * Constraints:
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

#include "leetcode.h"
// @lc code=start
class Solution {
public:
  int findKthLargest(vector<int> &nums, int k) {
    return findKthLargestHeap(nums, k);
  }
  int findKthLargesSort(vector<int> &nums, int k) {
    sort(nums.begin(), nums.end());
    return nums[nums.size() - k];
  }

  int findKthLargestHeap(vector<int> &nums, int k) {
    make_heap(std::begin(nums),std::end(nums),std::less());

    for (int i = 0; i < k-1; i++) {
        pop_heap( std::begin(nums),std::end(nums),std::less());
        nums.pop_back();
    }
    pop_heap(std::begin(nums), std::end(nums), std::less());
    return nums.back();
  }
  int findKthLargestQuickSelect(vector<int> &nums, int k) {
    
  }
  };
  // @lc code=end
