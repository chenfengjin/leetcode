/*
 * @lc app=leetcode id=432 lang=golang
 *
 * [432] All O`one Data Structure
 */
package main

// @lc code=start
type Node struct {
	left  *Node
	right *Node
	val   int
}
type AllOne struct {
	head  *Node
	tail  *Node
	nodes map[string]*Node
}

func Constructor() AllOne {

}

func (this *AllOne) Inc(key string) {
	node, ok := this.nodes[key]
	if !ok {
		node = &Node{
			left:  this.tail,
			right: nil,
			val:   1,
		}
		this.tail = node
		this.nodes[key] = node
		return
	}
	node.val += 1
	for node.right != nil && node.val > node.right.val {
		// if node.right.right != nil {
		// 	node.right.right.left = node
		// }
		// node.left.right = node.right
		// node.right.left = node.left
		// node.left
	}

}

func (this *AllOne) Dec(key string) {

}

func (this *AllOne) GetMaxKey() string {

}

func (this *AllOne) GetMinKey() string {

}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
// @lc code=end
