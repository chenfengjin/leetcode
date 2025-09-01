/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
 *
 * https://leetcode.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (46.80%)
 * Likes:    34586
 * Dislikes: 6922
 * Total Accepted:    6.1M
 * Total Submissions: 13.1M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * You are given two non-empty linked lists representing two non-negative
 * integers. The digits are stored in reverse order, and each of their nodes
 * contains a single digit. Add the two numbers and return the sum as a linked
 * list.
 * 
 * You may assume the two numbers do not contain any leading zero, except the
 * number 0 itself.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: l1 = [2,4,3], l2 = [5,6,4]
 * Output: [7,0,8]
 * Explanation: 342 + 465 = 807.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: l1 = [0], l2 = [0]
 * Output: [0]
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * Output: [8,9,9,9,0,0,0,1]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * The number of nodes in each linked list is in the range [1, 100].
 * 0 <= Node.val <= 9
 * It is guaranteed that the list represents a number that does not have
 * leading zeros.
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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int carry=0;
        ListNode *dummy = new ListNode();
        ListNode *cur = dummy;
        while (l1 && l2)
        {
            int val = (l1->val + l2->val+carry) % 10;
            carry = (l1->val + l2->val + carry) / 10;
            cur->next = new ListNode(val);
            cur = cur->next;
            l1 = l1->next;
            l2 = l2->next;
        }
        if(l2){
            l1 = l2;
        }
        while (l1){
            int val = (l1->val + carry) % 10;
            carry = (l1->val + carry) / 10;
            cur->next = new ListNode(val);
            cur = cur->next;
            l1 = l1->next;
        }
        if (carry){
            cur->next = new ListNode(carry);
        }
        return dummy->next;
    };
};
// @lc code=end

//  * Example 1:
//  * 
//  * 
//  * Input: l1 = [2,4,3], l2 = [5,6,4]
//  * Output: [7,0,8]
//  * Explanation: 342 + 465 = 807.

TEST(AddTwoNumbers,Case1){
    vector<int> v1 = {2, 4, 3};
    vector<int> v2 = {5, 6, 4};

    ListNode *l1 = create_list_from_vector_int(v1);
    ListNode *l2 = create_list_from_vector_int(v2);
    vector<int> vwant = {7, 0, 8};
    ListNode *want = create_list_from_vector_int(vwant);
    ListNode *out = Solution().addTwoNumbers(l1, l2);
    print_link_list(out);
    ASSERT_TRUE(link_list_equal(want, out));
}