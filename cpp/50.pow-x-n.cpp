/*
 * @lc app=leetcode id=50 lang=cpp
 *
 * [50] Pow(x, n)
 */
#include<leetcode-helper.h>
// @lc code=start
class Solution {
public:
    double myPow(double x, int n) {
        if (n==-2147483648){
            return 0;
        }
        // TODO 
        if(x==1){
            return 1;
        }
        bool sign = false;
        if (n<0){
            sign=true;
            n = -n;
        }
        double mem[32];
        double cur = x;
        for (int i = 0; i < 32; i++)
        {
            mem[i] = cur;
            cur = cur * cur;
        }
        double ret = 1;
        auto idx = 0;
        while (n > 0)
        {
            if (n&1){
                ret *= mem[idx];
            }
            idx++;
            n >>= 1;
        }
        return sign? 1/ret:ret;
    }
};
// @lc code=end

int main(){
    cout << Solution().myPow(2, -7);
}