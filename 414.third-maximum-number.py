#
# @lc app=leetcode id=414 lang=python3
#
# [414] Third Maximum Number
#

# @lc code=start
class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        max_number = None
        second_maximum_number = None
        third_maximum_number = None
        for num in nums:
            if num in [max_number,second_maximum_number,third_maximum_number]:
                continue
            if max_number is None:
                max_number = num
                continue
            if num > max_number:
                max_number,num = num,max_number
            if second_maximum_number is None:
                second_maximum_number = num
                continue
            if num > second_maximum_number:
                second_maximum_number,num = num,second_maximum_number
            if third_maximum_number is None:
                third_maximum_number = num
                continue
            if  num > third_maximum_number:
                third_maximum_number,num = num,third_maximum_number
        return third_maximum_number if third_maximum_number is not None else max_number
            


# @lc code=end


if __name__ == "__main__":
    pass