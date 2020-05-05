/*
 * @lc app=leetcode id=67 lang=java
 *
 * [67] Add Binary
 */

// @lc code=start
class Solution {
    public String addBinary(String a, String b) {
        if (a.length() > b.length()){
            String tmp = a;
            a=b;
            b=tmp;
        }
        for(int i = 0;i<a.length()-b.length();i++){
            b = "0" + b;
        }
        String sum = "";
//        sum.charAt()
        for(int i=a.length()-1;i>=0;i--){
            int c = (a.charAt(i)-'0' + b.charAt(i) -'0');
            int carry = c % 2;
            int sum = c/2;

        }
    }
}
// @lc code=end

