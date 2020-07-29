#
# @lc app=leetcode id=68 lang=python3
#
# [68] Text Justification
#
from typing import *
# @lc code=start
class Solution:
    def convertWordsToLine(self,words,maxWidth):

        num_charaters = sum([len(word) for word in words])
        num_words = len(words)

        num_spaces = maxWidth - num_charaters
        if num_words == 1:
            return words[0]+" "*num_spaces
        space_count = num_spaces // (num_words-1)
        wide_count = num_spaces % (num_words-1)
        left = (' '* (space_count + 1)).join(words[:wide_count])
        right =  (' '*space_count).join(words[wide_count:])
        return  left+  ' '* (space_count+1) + right if left else right

    def fullJustify(self, words: List[str], maxWidth: int) -> List[str]:
        #  all words are less than  16 characters
        left = 0
        right = 0
        total = 0
        result = []
        while right < len(words):
            total = -1
            while right  < len(words) and total +1 + len(words[right]) <= maxWidth:
                total += len(words[right])
                right += 1
                total += 1
            if right >= len(words):
                result.append(self.convertWordsToLine([" ".join(words[left:])],maxWidth))
            else:
                result.append(self.convertWordsToLine(words[left:right], maxWidth))
            left = right
        return result

        
# @lc code=end

if __name__ == "__main__":
    words = ["This", "is", "an", "example", "of", "text", "justification."]
    print(Solution().fullJustify(words = words,maxWidth = 16))
    words = ["Science","is","what","we","understand","well","enough","to","explain",
    "to","a","computer.","Art","is","everything","else","we","do"]
    print(Solution().fullJustify(words = words,maxWidth = 20))

