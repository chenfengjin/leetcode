/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 */
package main

import "fmt"

// @lc code=start
type NumArray struct {
	bit  []int
	nums []int
}

func Constructor(nums []int) NumArray {
	bit := make([]int, len(nums)+1)
	numArray := NumArray{
		bit:  bit,
		nums: make([]int, len(nums)),
	}
	for idx, num := range nums {
		numArray.Update(idx, num)
	}
	return numArray
}

func lowBit(x int) int {
	return x & (-x)
}

func (this *NumArray) Update(index int, val int) {
	delta := val - this.nums[index]
	this.nums[index] = val
	for index+1 < len(this.bit) {
		this.bit[index+1] += delta
		index += lowBit(index + 1)
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.Prefixrange(right+1) - this.Prefixrange(left)
}
func (this *NumArray) Prefixrange(n int) int {
	ret := 0
	for n > 0 {
		ret += this.bit[n]
		n -= lowBit(n)
	}
	return ret
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
func main() {
	m := Constructor([]int{1, 3, 5})
	obj := &m
	fmt.Println(obj.SumRange(0, 2))
	obj.Update(1, 2)
	fmt.Println(obj.SumRange(0, 2))
}
