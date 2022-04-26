/*
 * @lc app=leetcode id=191 lang=cpp
 *
 * [191] Number of 1 Bits
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    int hammingWeight(uint32_t n) {
        auto ret = 0;
        while(n){
            n &= n - 1;
            ret++;
        }
        return ret;
    }
};
// @lc code=end

