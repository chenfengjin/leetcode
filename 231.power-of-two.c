/*
 * @lc app=leetcode id=231 lang=c
 *
 * [231] Power of Two
 */


bool isPowerOfTwo(int n){
      if(n<=0) return false;
        return !(n&(n-1));
}



