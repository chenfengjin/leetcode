#
# @lc app=leetcode id=279 lang=python3
#
# [279] Perfect Squares
#
import pdb
# @lc code= 
class Solution: #TODO 四平方和定理
    def numSquares(self, n: int) -> int:
        output = [1] #1
        squares = [1]
        for i in range(2,n + 1):
            if i == (len(squares)+1) ** 2:
                squares.append(i)
                output.append(1)
            else:
                min_step =  1 << 32
                for j in range(len(squares)):
                    try:
                        min_step = min(output[i -squares[j]-1 ] , min_step)
                    except:
                        pdb.set_trace()
                output.append(min_step + 1)
        return output[-1]

if __name__ == "__main__":
    print(Solution().numSquares(13))
# @lc code=end