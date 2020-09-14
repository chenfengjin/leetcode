#
# @lc app=leetcode id=137 lang=python3
#
# [137] Single Number II
#
from typing import *

# @lc code=start
# 总结一下
# 位运算的优先级
# 负数的处理
# 位数的考量
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        bits = [0]*32
        for num in nums:
            for i in range(32):
                bits[i] += (num >> i) & 0x01
        bits = [bit % 3 for bit in bits]
        
        total  = 0
        for i in range(32):
            total |= bits[i] << i
        # 1 左移，2 的次方，别混淆
        # 负数咋处理呢？
        # 方案1
            # 怎么右移的，就怎么左移回去
            # 好吧，不 work
        # 方案2
            # 按位取反再加1 -> 减1 再按位取反
            # 应该可以，太复杂，没尝试
        # 方案3
            # 完美,对补码的深入理解  
        if bits[-1]:
            total = total - (1<<32 )    # 位运算优先级居然比+/-还低，大哭
        return total 

if __name__ == "__main__":
    print(Solution().singleNumber([-2,-2,1,1,-3,1,-3,-3,-4,-2]))
    
        
# @lc code=end

