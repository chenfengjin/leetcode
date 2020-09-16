#
# @lc app=leetcode id=57 lang=python3
#
# [57] Insert Interval
#
from typing import *
# @lc code=start
class Solution: # Too slow ,should be optimized
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        if not intervals:
            return [newInterval]
        if [interval for interval in intervals if not interval[0] > newInterval[0] and not interval[1] < newInterval[1]]:
            return intervals
        overlay_intervals = [interval for interval in intervals if not interval[0] < newInterval[0] and not interval[1] > newInterval[1]]
        for overlay_interval in overlay_intervals:
            intervals.remove(overlay_interval)
        overlap_intervals = [interval for interval in intervals if 
        (not interval[0] > newInterval[1] and not interval[0] < newInterval[0])   
        or (not interval[1] < newInterval[0] and not interval[1] > newInterval[1])]
        for overlap_interval in overlap_intervals:
            intervals.remove(overlap_interval)
        overlap_intervals.append(newInterval)
        left = min(i[0] for i in overlap_intervals)
        right = max(i[1] for i in overlap_intervals)
        intervals.append([left,right])
        print(overlap_intervals)
        intervals.sort(key= lambda x:x[0])
        return intervals
            
        
# @lc code=end

if __name__ == "__main__":
    Solution().insert([[1,3],[6,9]],[2,5])
