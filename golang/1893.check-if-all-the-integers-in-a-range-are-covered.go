/*
 * @lc app=leetcode id=1893 lang=golang
 *
 * [1893] Check if All the Integers in a Range Are Covered
 */

// @lc code=start

type IntervalSlice [][]int

func (intervals IntervalSlice) Len() int {
	return len(intervals)

}
func (interval IntervalSlice) Less(i, j int) bool {
	return interval[i][0] < interval[j][0]
}

func (intervals *IntervalSlice) Swap(i, j int) {
	(*intervals)[i], (*intervals)[j] = (*intervals)[j], (*intervals)[i]
}
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	intervalSlice := IntervalSlice(intervals)
	sort.Sort(&intervalSlice)

	ret := [][]int{}
	cur := intervalSlice[0]
	for _, interval := range intervalSlice[1:] {
		if interval[0] <= cur[1] {
			if interval[1] > cur[1] {
				cur[1] = interval[1]
			}
		} else {
			ret = append(ret, cur)
			cur = interval
		}
	}
	ret = append(ret, cur)
	return ret

}
func isCovered(ranges [][]int, left int, right int) bool {
	ranges = merge(ranges)
	for i := 0; i < len(ranges); i++ {
		if ranges[i+1][0] > left {
			break
		}
	}
	return ranges[i][1] > right
}

// @lc code=end

