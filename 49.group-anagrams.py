#
# @lc app=leetcode id=49 lang=python3
#
# [49] Group Anagrams
#

# @lc code=start
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = {}
        for word in strs:
            key =  "".join(sorted(word))
            if not key  in groups:
                groups[key] =[]
            groups[key].append(word)
        return [value for _,value in groups.items()] 
# @lc code=end

