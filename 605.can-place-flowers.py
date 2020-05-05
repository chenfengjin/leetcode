#
# @lc app=leetcode id=605 lang=python3
#
# [605] Can Place Flowers
#

# @lc code=start
class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        flowerbed.append(0)
        flowerbed.insert(0,0)
        for i in range(1,len(flowerbed)-1):
            if not flowerbed[i] and not flowerbed[i-1] and not flowerbed[i+1]:
                flowerbed[i] = 1
                n -= 1
        return n <= 0
            
            
        
# @lc code=end

