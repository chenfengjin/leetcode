#
# @lc app=leetcode.cn id=28 lang=python3
#
# [28] 实现 strStr()
#

# @lc code=start
# 这个题目有一些比较好的东西，KMP/BM/Horspoll/Sunday etc.
class Solution:
    # TODO 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
    # 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
    def strStr(self, haystack: str, needle: str) -> int:
        # pythonic
        return haystack.find(needle)
# @lc code=end

