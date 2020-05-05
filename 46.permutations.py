#
# @lc app=leetcode id=46 lang=python3
#
# [46] Permutations
#
<<<<<<< HEAD

# @lc code=start
class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        pass
# @lc code=end

=======
from typing import *
import copy
# @lc code=start
class Solution:
    def __init__(self):
        self.out = []
        self.res = []

    def permute(self, nums: List[int]) -> List[List[int]]:
        self.length = len(nums) 
        self.visited = [False] *self.length
        self.helper(nums)
        return self.res
    def helper(self,nums):
        if len(self.out) ==self.length:
            self.res.append(copy.copy(self.out))
            return
        for index,num in enumerate(nums):   
            if self.visited[index]:
                continue
            self.out.append(num)
            self.visited[index] = True
            self.helper(nums)
            self.out.pop()
            self.visited[index] = False
    
    # def permute(self, nums: List[int]) -> List[List[int]]:
    #     prefixes = [[]]
    #     out = []
    #     for num in nums:
    #         new_prefix = [prefix + [num] for prefix in prefixes]
    #         prefixes.extend([i for i in new_prefix if len(i)<len(nums)])
    #         out.extend([i for i in new_prefix if len(i)==len(nums)])
    #         print('profixed:{},out:{}'.format(prefixes,out))  
    #     return out          
        
# @lc code=end


if __name__ == "__main__":
    print(Solution().permute([1,2,3]))
>>>>>>> RFC:change mac
