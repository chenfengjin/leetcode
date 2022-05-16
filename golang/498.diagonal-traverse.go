/*
 * @lc app=leetcode id=498 lang=golang
 *
 * [498] Diagonal Traverse
 */
package main

import "fmt"

// 弄鬼细节题，核心包括两个部分
// 1. 将矩阵放到第一象限，就变成可 X+Y=K
// 2. 直接通用遍历，如果超出界限就跳过
// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}
	rows := len(mat)
	columns := len(mat[0])
	direction := 1
	ret := []int{}
	// 注意这里需要减1
	for i := 0; i < rows+columns-1; i++ {
		for j := 0; j <= i; j++ {
			if direction == 1 {
				if j >= columns || i-j >= rows {
					continue
				}
				ret = append(ret, mat[i-j][j])
			} else {
				if j >= rows || i-j >= columns {
					continue
				}
				ret = append(ret, mat[j][i-j])
			}
		}
		direction ^= 1
	}
	return ret
}

// @lc code=end
func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}))

}
