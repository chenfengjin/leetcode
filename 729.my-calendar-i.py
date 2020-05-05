#
# @lc app=leetcode id=729 lang=python3
#
# [729] My Calendar I
#

# @lc code=start
class MyCalendar: #TODO 这个太慢了
    def __init__(self):
        self.starts = []
        self.ends = []
        

    def book(self, start: int, end: int) -> bool:
        self.starts.append(start)
        self.starts.sort()
        self.ends.append(end)
        self.ends.sort()
        for i in range(len(self.ends)-1):
            if self.ends[i] > self.starts[i+1]:
                self.starts.remove(start)
                self.ends.remove(end)
                return False
        return True

# Your MyCalendar object will be instantiated and called as such:
# obj = MyCalendar()
# param_1 = obj.book(start,end)
# @lc code=end

