#
# @lc app=leetcode id=232 lang=python3
#
# [232] Implement Queue using Stacks
#

# @lc code=start
class Stack():
    def __init__(self):
        self.elements = []

    def push(self,e):
        self.elements.append(e)

    def pop(self):
        return self.elements.pop()

    def __len__(self):
        return len(self.elements)

    def empty(self):
        return not self.elements

    def top(self):
        return self.elements[-1]


class MyQueue:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.left = Stack()
        self.right = Stack()
        

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        self.right.push(x)
        

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        if self.left.empty() and not self.right.empty():
            while not self.right.empty():
                self.left.push(self.right.pop())

        return self.left.pop()

    def peek(self) -> int:
        """
        Get the front element.
        """
        if self.left.empty() and  not self.right.empty():
            while not self.right.empty():
                self.left.push(self.right.pop())

        return self.left.top()

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return self.left.empty() and self.right.empty()


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
# @lc code=end

