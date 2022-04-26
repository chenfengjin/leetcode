/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
 */

/**
 * Definition for singly-linked list.
 */
  struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
  };
 /* */

// @lc code=start

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        auto carry = 0;
        auto* ret =  new ListNode();
        auto cur = ret;
        while (l1 && l2)
        {
            auto total = l1->val + l2->val + carry;
            cur->next = new ListNode(total % 10);
            carry = total / 10;
            l1 = l1->next;
            l2 = l2->next;
            cur = cur->next;
        }
        if (!l1){
            l1 = l2;
        }
        while(l1){
            auto total = l1->val  + carry;
            cur->next = new ListNode(total % 10);
            carry = total / 10;
            l1 = l1->next;
            cur = cur->next;
        }
        if (carry){
            cur->next = new ListNode(carry);
        }
        return ret->next;
    }
};
// @lc code=end

