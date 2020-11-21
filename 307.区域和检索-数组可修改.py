#
# @lc app=leetcode.cn id=307 lang=python3
#
# [307] 区域和检索 - 数组可修改
#

# @lc code=start
class NumArray:
    def __init__(self, nums: List[int]): 
        self.tree = [0]*len(nums)*2
        self.n = len(nums)
        self.tree[n+1:] = nums
        for i in range(n,0,-1):
            # TODO 这里是 i*2+1 因为 下标从0 开始
            tree[i] = tree[i*2] + tree[i*2+1]

    def update(self, i: int, val: int) -> None:
        i = self.n + i
        nums[i] = val
        while i // 2 : 
            i//=2 #TODO i==0?
            nums[i] = nums[i*2] + nums[i*2+1]
    def sumRange(self, i: int, j: int) -> int:
        total = 0
        while(i<j):
            if i&1:
                total += self.tree[i]
                i+=1
            if not j&1:
                total += self.tree[j]
                j-=1
            j//=2
            j//=2
        return total

# Your NumArray object will be instantiated and called as such:
# obj = NumArray(nums)
# obj.update(i,val)
# param_2 = obj.sumRange(i,j)
# @lc code=end

