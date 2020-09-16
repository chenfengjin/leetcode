#
# @lc app=leetcode id=204 lang=python3
#
# [204] Count Primes
#

# @lc code=start
class Solution:
    def countPrimes(self, n: int) -> int:
        if n < 2:
            return 0
        nums = [1] * n
        nums[0] = 0
        nums[1] = 0

        def set_flag(i):
            index = i * 2
            while index < n:
                nums[index] = 0
                index += i

        for i in range(2,n):
            if nums[i]:
                set_flag(i)
        # print(nums)
        return len([num for num in nums if num])
                
        
# @lc code=end

if __name__ =="__main__":
    print(Solution().countPrimes(2))