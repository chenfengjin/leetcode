#
# @lc app=leetcode id=165 lang=python3
#
# [165] Compare Version Numbers
#

# @lc code=start
class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        version1 = [int(i) for i in version1.split(".")]
        version2 = [int(i) for i in version2.split(".")]
        if len(version1) > len(version2):
            version2.extend([0]*(len(version1) - len(version2)))
        if len(version1) < len(version2):
            version1.extend([0]*(len(version2)-len(version1)))    
        for v1,v2 in zip(version1,version2):
            if v1 > v2:
                return 1
            if v1 < v2:
                return -1
        return 0    
# @lc code=end

