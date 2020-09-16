#
# @lc app=leetcode id=345 lang=python3
#
# [345] Reverse Vowels of a String
#

# @lc code=start
class Solution:
    def reverseVowels(self, s: str) -> str:
        # Not good
        s = list(s)
        vowel_list = ["a","e","i","o","u","A","E","I","O","U"]
        vowels = [i for i in s if i in vowel_list]
        for (index,c) in enumerate(s):
            if c in vowel_list:
                s[index] = vowels.pop()
        return "".join(s)


        
# @lc code=end


if __name__ == "__main__":
    assert(Solution().reverseVowels("leetcode")=="leotcede")