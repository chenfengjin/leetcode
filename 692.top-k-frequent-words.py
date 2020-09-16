#
# @lc app=leetcode id=692 lang=python3
#
# [692] Top K Frequent Words
#

import heapq
from typing import *
# @lc code=start
# TODO 自己实现大顶堆
class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        word_count_map = {}
        count_word_map = {}
        counts_high_to_low = []
        result = []
        for word in words:
            if not word in word_count_map:
                word_count_map[word] = 0
            word_count_map[word] += 1
        for word,count in word_count_map.items():
            if not count in count_word_map:
                count_word_map[count] = []
            count_word_map[count].append(word)
        counts = list(count_word_map.keys())
        heapq.heapify(counts)
        while counts:
            counts_high_to_low.append(heapq.heappop(counts))
        print(counts_high_to_low)
        while len(result) < k:
            print(counts_high_to_low)
            count = counts_high_to_low.pop()
            words = count_word_map[count]
            words.sort()
            for word in words:
                result.append(word)
        return result[0:k]            


        
        
# @lc code=end

if __name__ == "__main__":
    print(Solution().topKFrequent(["i", "love", "leetcode", "i", "love", "coding"],2))