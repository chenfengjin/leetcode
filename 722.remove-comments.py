#
# @lc app=leetcode id=722 lang=python3
#
# [722] Remove Comments
#

# @lc code=start
class Solution:
    def removeComments(self, source: List[str]) -> List[str]:
        result = []
        block_comment = False
        partial_line = ""
        index = 0
        flag = False
        while index < len(source):
            source_line = source[index]
            while True:
                if block_comment:
                    if flag:
                        source_line = source_line[source_line.index("/*")+2:]
                        flag = False
                    if  "*/" in source_line:
                        source_line  = source_line[source_line.index("*/")+2:]
                        block_comment = False
                        if partial_line+source_line:
                            result.append(partial_line+source_line)
                    index += 1
                    continue

                source_line = source_line.split("//")[0]
                if  "/*" in source_line:
                    partial_line = source_line.split("/*")[0]
                    flag = True
                    block_comment = True
                else:
                    if source_line:
                        result.append(source_line)
                    index += 1
        return result       
        
# @lc code=end

