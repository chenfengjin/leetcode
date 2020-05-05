#
# @lc app=leetcode id=336 lang=python3
#
# [336] Palindrome Pairs
#

# @lc code=start
class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        words_map = {}
        result = []
        for (index,word) in enumerate(words_map):
            words_map[word] = index
        for (index,word) in enumerate(words):
            if word[::-1] in words and not words_map[word[::-1]] == index:
                result.append(index,words_map[word])
            if len(word) > 1:
                for i in range(1,len(word)):
                    if word[len(word):i:-1] in words_map:
                        result.append([index,words_map[word]])
        return result

        
# @lc code=end

