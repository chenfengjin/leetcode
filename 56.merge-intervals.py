#
# @lc app=leetcode id=56 lang=python3
#
# [56] Merge Intervals
#

# @lc code=start
List=list
class Solution:
    # def merge(self, intervals: List[List[int]]) -> List[List[int]]:
    def merge(self, intervals):
        if intervals==[]:
            return []
        result=[]
        intervals.sort(key=lambda x: x[0])
        current=intervals[0]
        for interval in intervals[1:]:

            if interval[0] > current[1]:
                result.append(current)
                current=interval
                continue
            if interval[1] > current[1]:
                current[1] = interval[1] 
        
        result.append(current)
        return result
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().merge(intervals=[[1,3],[2,6],[8,10],[15,18]]))