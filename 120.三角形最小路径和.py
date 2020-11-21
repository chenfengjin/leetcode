#
# @lc app=leetcode.cn id=120 lang=python3
#
# [120] 三角形最小路径和
#

# @lc code=start
class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        # Yeah, I make it with O(1) extrac space  
        for row_index,row in enumerate(triangle):
            if row_index == 0:
                continue
            for column_index,column in enumerate(row):
                if column_index == 0:
                    triangle[row_index][0] += triangle[row_index-1][0]
                elif column_index == row_index:
                    triangle[row_index][column_index] += triangle[row_index-1][column_index-1]
                else:
                    triangle[row_index][column_index] += min(triangle[row_index-1][column_index],triangle[row_index-1][column_index-1])
        return min(triangle[-1])
# @lc code=end

