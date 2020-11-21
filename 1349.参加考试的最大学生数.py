#
# @lc app=leetcode.cn id=1349 lang=python3
#
# [1349] 参加考试的最大学生数
#

# @lc code=start
class Solution:
    def maxStudents(self, seats: List[List[str]]) -> int:
        def bit_count(x:int):
            if x == 0:
                return 0
            cnt = 0
            while x:
                cnt += 1
                x = x & (x - 1)
            return cnt
        # 状态压缩DP
        # 预处理 seats
        n, m = len(seats), len(seats[0])
        st = []
        f = [[0] * (1 << m) for _ in range(n)]
        for row in seats:
            state = 0
            # 计算这一行的初始状态 有空座是1
            for i in range(m - 1, -1, -1):
                if row[i] == '#':
                    state = state * 2
                else:
                    state = state * 2 + 1
            st.append(state)
        for i in range(n):
            for j in range(1 << m):
                # 如果这个状态可以坐下：
                # j & st[i] == j 判断j的状态下能否坐下第i横排
                # (j & (j >> 1) == 0) 判断j模式左右是否没人
                if j & st[i] == j and j & (j >> 1) == 0:
                    if i == 0:
                        f[i][j] = bit_count(j)
                    else:
                        # 枚举上一行的状态
                        for k in range(1 << m):
                            if k & (j >> 1) == 0 and (k >> 1) & j == 0:
                                f[i][j] = max(f[i - 1][k] + bit_count(j), f[i][j])
        return max(f[n - 1])
        
# @lc code=end

