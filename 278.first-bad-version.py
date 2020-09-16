#
# @lc app=leetcode id=278 lang=python3
#
# [278] First Bad Version
#

# @lc code=start
# The isBadVersion API is already defined for you.
# @param version, an integer
# @return a bool
def isBadVersion(version):
    return [False,False,False,False,True,True][version]

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        left = 1
        right = n
        middle = 1
        while left < right:
            middle = (left+right) // 2
            if isBadVersion(middle):
                right = middle
            else:
                left = middle + 1
        return middle
            
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().firstBadVersion(5))
