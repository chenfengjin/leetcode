/*
 * @lc app=leetcode id=23 lang=cpp
 *
 * [23] Merge k Sorted Lists
 *
 * https://leetcode.com/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (57.48%)
 * Likes:    20694
 * Dislikes: 768
 * Total Accepted:    2.6M
 * Total Submissions: 4.6M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * You are given an array of k linked-lists lists, each linked-list is sorted
 * in ascending order.
 *
 * Merge all the linked-lists into one sorted linked-list and return it.
 *
 *
 * Example 1:
 *
 *
 * Input: lists = [[1,4,5],[1,3,4],[2,6]]
 * Output: [1,1,2,3,4,4,5,6]
 * Explanation: The linked-lists are:
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * merging them into one sorted linked list:
 * 1->1->2->3->4->4->5->6
 *
 *
 * Example 2:
 *
 *
 * Input: lists = []
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: lists = [[]]
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] is sorted in ascending order.
 * The sum of lists[i].length will not exceed 10^4.
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
  ListNode *mergeKLists(vector<ListNode *> &lists) {
    if (lists.size() == 0) {
      return nullptr;
    }
    // 二路归并性能还可以
    while (lists.size() > 1) {
      for (int i = 0; i < lists.size() >> 1; i++) {
        lists[i] = mergeTwoLists(lists[i], lists[lists.size() - 1 - i]);
      }
      auto pop_count = lists.size() >> 1;
      for (int i = 0; i < pop_count; i++) {
        lists.pop_back();
      }
    }
    return lists[0];
  }
  ListNode *slowMergeKLists(vector<ListNode *> &lists) {

    while (lists.size() > 1) {
      auto list1 = lists.back();
      lists.pop_back();
      auto list2 = lists.back();
      lists.pop_back();
      auto partial_merged = mergeTwoLists(list1, list2);

      lists.push_back(partial_merged);
    }
    return lists[0];
  }
  void erase_list(ListNode *list) {
    if (!list) {
      return;
    }
    erase_list(list->next);
    delete list;
  }

  ListNode *mergeTwoLists(ListNode *list1, ListNode *list2) {

    ListNode *dummy = new ListNode();
    ListNode *cur = dummy;
    while (list1 && list2) {
      if (list1->val < list2->val) {
        cur->next = list1;
        list1 = list1->next;
      } else {
        cur->next = list2;
        list2 = list2->next;
      }
      cur = cur->next;
    }
    list1 = list2 ? list2 : list1;
    while (list1) {
      cur->next = list1;

      list1 = list1->next;
      cur = cur->next;
    }

    return dummy->next;
  }
};
// @lc code=end
