#
# @lc app=leetcode id=166 lang=python3
#
# [166] Fraction to Recurring Decimal
#
from typing import *
# @lc code=start
# 通过一个哈希表记录 numerator 出现的次数
# 乱的一匹配，有空整理下
class Solution:
    def fractionToDecimal(self, numerator: int, denominator: int) -> str:
        sign  =  '-' if numerator > 0 and denominator < 0 or numerator < 0 and denominator > 0   else "" 
        numerator = abs(numerator)
        denominator = abs(denominator)

        h = set()
        integar = numerator // denominator
        remain = numerator % denominator
        float1 = []
        while  not remain in h:
            h.add(remain)
            float1.append(remain * 10 // denominator)
            remain  = remain * 10 % denominator 
        while len(float1) and float1[-1] == 0:
            float1.pop()
        float1 = [str(i) for i in float1]
        if not float1:
            return sign + str(integar)
        if not remain :
            return sign + "{}.{}".format(integar,"".join(float1))

        segment_length = 1
        old = remain
        while not remain * 10 % denominator ==  old:
            remain  = remain * 10 % denominator 
            segment_length += 1
        return sign + "{}.{}({})".format(str(integar),"".join(float1[:-1 * segment_length]),"".join(float1[-1 * segment_length:]))

# @lc code=end
if __name__ == "__main__":
    # print(Solution().fractionToDecimal(2,3))
    # print(Solution().fractionToDecimal(1,2))
    print(Solution().fractionToDecimal(2,1))
    # print(Solution().fractionToDecimal(4,333))
