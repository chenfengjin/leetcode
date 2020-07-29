#
# @lc app=leetcode id=233 lang=python3
#
# [233] Number of Digit One
#

# @lc code=start
class Solution:
    def countDigitOne(self, n: int) -> int:
        counts = [1,]
        res = 0
        bits = [int(i) for i in list(str(n))]
        count_map = [0]
        for i in  range(1,len(bits)):
            count_map.append(10**(i-1)+10*count_map[-1])
            # print(count_map)
            # count_map.append(10**(i-1)+9*count_map[-1])
        # count_map.pop(0)
        # count_map = [sum(count_map[:i+1]) for i in range(len(count_map))]
        count_map.reverse()
        # print(count_map)
        print(count_map)
        for (index,bit) in enumerate(bits):
            res += bit * count_map[index]
            if bit > 1:
                pass
            elif bit == 1:
                pass
            else:
                pass
        return res

        
        
# @lc code=end
if __name__ == "__main__":
    print(Solution().countDigitOne(54321))
