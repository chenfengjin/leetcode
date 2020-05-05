#
# @lc app=leetcode id=273 lang=python3
#
# [273] Integer to English Words
#

# @lc code=start
class Solution:
    def numberToWords(self, num: int) -> str:
        if num == 0:
            return "Zero"

        res =[]
        segment = 10 ** 3
        ones = num % segment
        num = num // segment

        thousands = num % segment
        num = num // segment

        millions = num % segment
        num = num // segment

        billions = num % segment
        num = num // segment    

        if billions:
            res.extend(self.transform_segment(billions,"billion"))
        if millions:
            res.extend(self.transform_segment(millions,"million"))
        if thousands:
            res.extend(self.transform_segment(thousands,"thousand"))
        if ones:
            res.extend(self.transform_segment(ones))
        return " ".join([i.capitalize() for i in res]).strip().replace("  "," ") # TODO so ugly

    def transform_segment(self,num,suffix = None): # deal with zero
        res = []
        if  num >= 100:
            res.extend([
            {
                1:"one",
                2:"two",
                3:"three",
                4:"four",
                5:"five",
                6:"six",
                7:"seven",
                8:"eight",
                9:"nine",
            }.get(num//100),"hundred"])
            num = num % 100
        if num >= 20:
            res.append(
                {
                    2:"twenty",
                    3:"thirty",
                    4:"forty",
                    5:"fifty",
                    6:"sixty",
                    7:"seventy",
                    8:"eighty", 
                    9:"ninety"
                    }.get(num//10)
            )
            num = num % 10

        res.append({
            0:"",
            1:"one",
            2:"two",
            3:"three",
            4:"four",
            5:"five",
            6:"six",
            7:"seven",
            8:"eight",
            9:"nine",
            10:"ten",
            11:"eleven",
            12:"twelve",
            13:"thirteen",
            14:"fourteen",
            15:"fifteen",
            16:"sixteen",
            17:"seventeen",
            18:"eighteen",
            19:"nineteen",
        }.get(num))
      
        if suffix is not None:
            res.append(suffix)
        return res 
# @lc code=end




if __name__ == "__main__":
    # print(Solution().transform_segment(123,""))
    # print(Solution().transform_segment(10,"million"))
    print(Solution().transform_segment(20))