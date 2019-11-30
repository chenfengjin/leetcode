#
# @lc app=leetcode id=71 lang=python3
#
# [71] Simplify Path
#

# @lc code=start
class Solution:
    def simplifyPath(self, path: str) -> str:
        result=[]
        direntries =  path.split("/")
        while "" in direntries:
            direntries.remove("")
        while "." in direntries:
            direntries.remove(".")
        for direntry in direntries:
            if direntry == "..": 
                if len(result):
                    result.pop()
            else:
                result.append(direntry)
        return "/"+"/".join(result)
        
# @lc code=end

if __name__ =="__main__":
        
    print(Solution().simplifyPath("/a/./b/../../c/"))
