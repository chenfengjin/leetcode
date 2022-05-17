/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU Cache
 */
package main

// @lc code=start
type DListNode struct {
	key  int
	val  int
	prev *DListNode
	next *DListNode
}
type LRUCache struct {
	head     *DListNode
	tail     *DListNode
	data     map[int]*DListNode
	capacity int
}

func Constructor(capacity int) LRUCache {
	head := &DListNode{}
	tail := &DListNode{}
	head.next = tail
	tail.next = head
	cache := LRUCache{
		head:     head,
		tail:     tail,
		data:     map[int]*DListNode{},
		capacity: capacity,
	}
	return cache
}

func (this *LRUCache) Get(key int) int {
	item, ok := this.data[key]
	if !ok {
		return -1
	}
	item.next.prev = item.prev
	item.prev.next = item.next

	item.next = this.head.next
	this.head.next.prev = item

	this.head.next = item
	item.prev = this.head
	return item.val
}

func (this *LRUCache) Put(key int, value int) {
	item, ok := this.data[key]
	if ok {
		item.val = value
		this.moveToFront(item)
		return
	}
	if len(this.data) == this.capacity {
		this.evict()
	}
	item = &DListNode{
		key: key,
		val: value,
	}
	item.next = this.head.next
	this.head.next.prev = item
	item.prev = this.head
	this.head.next = item
	this.data[key] = item
}
func (this *LRUCache) moveToFront(node *DListNode) {
	node.next.prev = node.prev
	node.prev.next = node.next

	node.next = this.head.next
	this.head.next.prev = node
	node.prev = this.head
	this.head.next = node

}
func (this *LRUCache) evict() {
	node := this.tail.prev
	delete(this.data, node.key)
	node.prev.next = this.tail
	this.tail.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Get(1)
}
