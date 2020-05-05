#
# @lc app=leetcode id=207 lang=python3
#
# [207] Course Schedule
#

from typing import *
# @lc code=start
class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        graph = [[0 for i in range(numCourses)]for j in range(numCourses)]
        courses_taken_num = 0
        for prerequisite in prerequisites:
            graph[prerequisite[0]][prerequisite[1]] = 1
        courses = list(range(numCourses))
        courses_can_take = [i for i in courses if sum(graph[i]) == 0 ]
        if not courses_can_take:
            return False
        while courses_can_take and courses_taken_num != numCourses:
            print(courses_can_take)
            courses_taken_num += 1 
            course = courses_can_take.pop()
            for i in range(numCourses):
                if graph[i][course] == 1:
                    graph[i][course] = 0
                    if sum(graph[i])  == 0:
                        courses_can_take.append(i)
        return True if courses_taken_num == numCourses else False          
 
if __name__ == "__main__":
    Solution().canFinish(2, [[1,0]])

# @lc code=end

