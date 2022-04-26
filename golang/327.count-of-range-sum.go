/*
 * @lc app=leetcode id=327 lang=golang
 *
 * [327] Count of Range Sum
 */
package main

import (
	"fmt"
	"sort"
)

// @lc code=start
func lowBit(n int) int {
	return n & -n
}

type BIT struct {
	elems []int
}

func newBIT(size int) *BIT {
	return &BIT{
		elems: make([]int, size),
	}
}

func (t *BIT) Add(x, val int) {
	t.elems[x] += val
	x = x + lowBit(x)
	// TODO
	for x < len(t.elems) {
		t.elems[x] += val
		x += lowBit(x)
	}

}
func (t *BIT) Query(x int) int {
	ret := 0
	for x >= 1 {
		ret += t.elems[x]
		x -= lowBit(x)
	}
	return ret

}
func (t *BIT) Range(low, high int) int {
	return t.Query(high) - t.Query(low)
}

// func(t*BIT)range(low,hgith int)int{

// }

func countRangeSum(nums []int, lower int, upper int) int {
	return 0
}

// @lc code=end

func main() {
	sort.Ints()

	t := newBIT(10)
	for i := 1; i < 10; i++ {
		t.Add(i, i)
	}
	for i := 1; i < 10; i++ {
		fmt.Println(t.Query(i))
	}
}
