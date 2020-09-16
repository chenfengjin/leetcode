/*
 * @lc app=leetcode id=69 lang=java
 *
 * [69] Sqrt(x)
 */

// @lc code=start
class Solution {
    public int mySqrt(int x) { // my first python program
        // pay attention to overflow
        if (x == 0 || x == 1) {
            return x;
        }
        if (x == 2) {
            return 1;
        }
        for (int i = 2; i < x; i++) {
            if (i * i == x) {
                return i;
            } else if (i * i > x || i * i < 0) {
                return i - 1;
            } else {
                continue;
            }
        }
        return -1;
    }
}
// @lc code=end
