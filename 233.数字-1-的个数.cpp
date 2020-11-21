/*
 * @lc app=leetcode.cn id=233 lang=cpp
 *
 * [233] 数字 1 的个数
 */

// @lc code=start
class Solution {
    #define mset(x,v) memset(x,(v),sizeof(x))
    #define rep(i,a,n) for(int i=(a);i<=(n);i++)
    #define ll long long
    #define maxn 15
    ll a[maxn],pw[maxn]={1},dp[maxn];
    ll dfs(int pos,int cnt,bool limit){
        if(!pos) return cnt;
        if(~dp[pos]&&!limit) return dp[pos]+cnt*pw[pos];
        ll ret=0,bound=limit?a[pos]:9;
        rep(i,0,bound) ret+=dfs(pos-1,cnt+(i==1),i==bound&&limit);
        if(!limit) dp[pos]=ret;
        return ret;
    }
    ll solve(ll x){
        int len=0;
        while(x) a[++len]=x%10,x/=10;
        mset(dp,-1);
        return dfs(len,0,1);
    }
public:
    int countDigitOne(int n) {
        rep(i,1,maxn-1) pw[i]=pw[i-1]*10;
        return solve(n);
    }
};

// @lc code=end

