/*
 * @lc app=leetcode id=155 lang=cpp
 *
 * [155] Min Stack
 */
#include<leetcode-helper.h>
#include<cstdint>
// @lc code=start
class MinStack {
public:
    MinStack() {
        min = INT32_MAX;;
    }

    void push(int val) {
        normalStack.push(val);
        if (min>val){
            min = val;
        }
        minStack.push(min);
    }

    void pop() {
        auto val = normalStack.top();
        // FIXME here is the cause of error
        minStack.pop();
        if (min<val){
            min = val;
        }
        normalStack.pop();
    }

    int top() {
        return normalStack.top();
    }

    int getMin() {
        return minStack.top();
    }

private:
    stack<int> minStack;
    int min;
    stack<int> normalStack;
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
// @lc code=end

