#
# @lc app=leetcode id=859 lang=python3
#
# [859] Buddy Strings
#

# @lc code=start
class Solution:
    def buddyStrings(self, A: str, B: str) -> bool:
        if not len(A) == len(B):
            return False
        if A==B:
            return True
        length = len(A)
        diff = [ord(A[i])-ord(B[i]) for i in range(length)]
        print(diff)
        return sum(diff) == 0 and len([i for i in diff if not i ==0]) ==2 
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().buddyStrings("ab","ab"))