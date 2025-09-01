/*
 * @lc app=leetcode id=24 lang=cpp
 *
 * [24] Swap Nodes in Pairs
 *
 * https://leetcode.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (67.77%)
 * Likes:    12706
 * Dislikes: 490
 * Total Accepted:    1.7M
 * Total Submissions: 2.5M
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a linked list, swap every two adjacent nodes and return its head. You
 * must solve the problem without modifying the values in the list's nodes
 * (i.e., only nodes themselves may be changed.)
 *
 *
 * Example 1:
 *
 *
 * Input: head = [1,2,3,4]
 *
 * Output: [2,1,4,3]
 *
 * Explanation:
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: head = []
 *
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: head = [1]
 *
 * Output: [1]
 *
 *
 * Example 4:
 *
 *
 * Input: head = [1,2,3]
 *
 * Output: [2,1,3]
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the list is in the range [0, 100].
 * 0 <= Node.val <= 100
 *
 *
 */

#include "leetcode.h"
// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
  ListNode *swapPairs(ListNode *head) {
    auto dummy = new ListNode(0, head);
    auto cur = dummy;
    // dummy 1 2 3 4
    // cur  dummy temp1
    // cur ->next 1 temp2
    // cur->next-> next 2   temp3
    //                      3 temp4
    // dummy 1 2 3 4 5
    // dummy 1
    // dummy 1 2
    // cur->1 1->2 2->3

    while (cur && cur->next && cur->next->next) {
      auto temp1 = cur;
      auto temp2 = cur->next;
      auto temp3 = cur->next->next;
      auto temp4 = temp3->next;
      temp2->next = temp4;
      temp3->next = temp2;
      temp1->next = temp3;
      cur = temp2;
    }

    return dummy->next;
  }
};
// @lc code=end
