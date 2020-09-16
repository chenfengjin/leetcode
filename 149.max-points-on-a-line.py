#
# @lc app=leetcode id=149 lang=python3
#
# [149] Max Points on a Line
#
from typing import *

# @lc code=start
# Too difficult, should be simplified
class Solution:
    def onOneLine(self,point_i, point_j,point_k):
        x1, y1 = point_i
        x2, y2 = point_j
        x3, y3 = point_k
        return x1*y2+ x2*y3+x3*y1 - x2*y1- x3* y2 -x1*y3 == 0
        # x1 y1 1
        # x2 y2 1
        # x3 y3 1
    def maxPoints(self, points: List[List[int]]) -> int:
        
        point_count = len(points)
        if point_count < 3:
            return point_count
        

        points = sorted(points)
        new_points = [points[0]]
        m = 1
        points_duplication_count = [1]
        while m  < len(points):
            if not points[m] == points[m-1]:
                new_points.append(points[m])
                points_duplication_count.append(1)
            else:
                points_duplication_count[-1] += 1
            m += 1 
        

        points = new_points

        if len(points) == 1:
            return points_duplication_count[-1]
        if len(points) == 2:
            return sum(points_duplication_count)

        max_point_on_a_count = 0
        current_point_count_in_a_line = 2
        i, j, k = 0, 0, 0
        point_count = len(points)

        while i < point_count:
            j = i + 1
            while j < point_count:
                k = j + 1
                current_point_count_in_a_line = points_duplication_count[i] + points_duplication_count[j]
                while k < point_count:
                    if self.onOneLine(points[i],points[j],points[k]):
                        current_point_count_in_a_line += points_duplication_count[k]
                    k += 1
                if current_point_count_in_a_line > max_point_on_a_count:
                    max_point_on_a_count = current_point_count_in_a_line
                j += 1
            i += 1

        return max_point_on_a_count

        
# @lc code=end

if __name__ == "__main__":
    # print(Solution().onOneLine([1,3],[2,6],[3,9]))
    # print(Solution().onOneLine([1,3],[2,6],[3,8]))
    # print(Solution().maxPoints([[1,1],[1,1],[0,0],[3,4],[4,5],[5,6],[7,8],[8,9]]))  
    print(Solution().maxPoints([[0,0],[1,1],[0,0]]))