#
# @lc app=leetcode id=210 lang=python3
#
# [210] Course Schedule II
#
from typing import *
# @lc code=start
class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        graph = [[0 for i in range(numCourses)]for j in range(numCourses)]
        courses_taken = []
        for prerequisite in prerequisites:
            graph[prerequisite[0]][prerequisite[1]] = 1
        courses = list(range(numCourses))
        courses_can_take = [i for i in courses if sum(graph[i]) == 0 ]
        if not courses_can_take:
            return []
        while courses_can_take and len(courses_taken) != numCourses:
            # print(courses_can_take)
            course = courses_can_take.pop()
            courses_taken.append(course)
            for i in range(numCourses):
                if graph[i][course] == 1:
                    graph[i][course] = 0
                    if sum(graph[i])  == 0:
                        courses_can_take.append(i)
        return courses_taken if len(courses_taken) == numCourses else []
# @lc code=end

if __name__ == "__main__":
    print(Solution().findOrder(4, [[1,0],[2,0],[3,1],[3,2]]))

