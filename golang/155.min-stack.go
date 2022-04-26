/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */
package main

// @lc code=start
type MinStack struct {
	real []int
	min  []int
	size int
}

func Constructor() MinStack {
	return MinStack{
		real: []int{},
		min:  []int{},
		size: 0,
	}
}

func (this *MinStack) Push(val int) {
	this.real = append(this.real, val)
	min := val
	if this.size != 0 && this.min[this.size-1] < min {
		min = this.min[this.size-1]
	}
	this.min = append(this.min, min)
	this.size += 1
}

func (this *MinStack) Pop() {
	this.min = this.min[0 : this.size-1]
	this.real = this.real[0 : this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.real[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
