#
# @lc app=leetcode id=567 lang=python3
#
# [567] Permutation in String
#
import pdb
# @lc code=start
class Solution:

    def checkInclusion(self, s1: str, s2: str) -> bool:
        if len(s1) > len(s2):
            return False
        letter_count_s1 = [0]* 26
        letter_count_s2 = [0]* 26
        
        for index  in range(len(s1)):
            letter_count_s1[ord(s1[index])-ord('a')] += 1
            letter_count_s2[ord(s2[index])-ord('a')] += 1

        left = 0
        right = len(s1)
        while right < len(s2) :
            if letter_count_s1 == letter_count_s2:
                return True
            letter_count_s2[ord(s2[left]) - ord('a')] -= 1
            letter_count_s2[ord(s2[right]) -ord('a')] += 1
            left += 1
            right += 1
        if letter_count_s1 == letter_count_s2:
            return True
        return False
    
        
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().checkInclusion("adc","dcda"))
    # print(Solution().checkInclusion("ab","eidboaoo"))