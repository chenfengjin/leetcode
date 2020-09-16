#
# @lc app=leetcode id=87 lang=python3
#
# [87] Scramble String
#

# @lc code=start
class Solution:
    def isScramble(self, s1: str, s2: str) -> bool:
        if not len(s1) == len(s2):
            return False
        return self.helper(s1,s2)
    
    def helper(self,s1,s2):
        if s1 == s2:
            print("s1:{},s2:{}".format(s1,s2))
            return True
        for i in range(len(s1)-1):
            if self.helper(s1[:i],s2[:i]):
                return self.helper(s1[i+1:],s2[i+1:])
        return False

        
# @lc code=end

if __name__ == "__main__":
    print(Solution().isScramble("abcde","caebd"))