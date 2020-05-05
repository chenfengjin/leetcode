#
# @lc app=leetcode id=44 lang=python3
#
# [44] Wildcard Matching
#

# @lc code=start
class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        while "**" in p:
            p=p.replace('**','*')
        # print("s:{},p:{}".format(s,p))
        return self.helper(s,p)
    def helper(self,s,p):
        # print(len(p))
        if not s and not p.replace('*',''):
            return True
        if s and not p:
            return False
        if not s and p:
            return False
        if p[0] == '?':
            return self.helper(s[1:],p[1:])
        elif p[0] == '*':
            return self.helper(s[1:],p) or self.helper(s[1:],p[1:]) or self.helper(s,p[1:]) 
        elif p[0] == s[0]:
            return self.helper(s[1:],p[1:])
        else:
            return False
        
# @lc code=end

if __name__ == "__main__":
    # print(Solution().isMatch("aa","a"))
    s="abbaabbbbababaababababbabbbaaaabbbbaaabbbabaabbbbbabbbbabbabbaaabaaaabbbbbbaaabbabbbbababbbaaabbabbabb"
    p="***b**a*a*b***b*a*b*bbb**baa*bba**b**bb***b*a*aab*a**"

    s="bbbbbbbabbaabbabbbbaaabbabbabaaabbababbbabbbabaaabaab"
    p="b*b*ab**ba*b**b***bba"

    s="babbbaabbaaaaabbababaaaabbbbbbbbbbabbaaaabbababbabaa"
    p="**a****a**b***ab***a*bab"

    s="bbaaaaaabbbbbabbabbaabbbbababaaabaabbababbbabbababbaba"
    p="b*b*ba**a*aaa*a*b**bbaa"

    print(Solution().isMatch(s,p))