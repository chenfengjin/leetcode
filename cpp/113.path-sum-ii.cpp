/*
 * @lc app=leetcode id=113 lang=cpp
 *
 * [113] Path Sum II
 *
 * https://leetcode.com/problems/path-sum-ii/description/
 *
 * algorithms
 * Medium (60.92%)
 * Likes:    8417
 * Dislikes: 165
 * Total Accepted:    1.1M
 * Total Submissions: 1.7M
 * Testcase Example:  '[5,4,8,11,null,13,4,7,2,null,null,5,1]\n22'
 *
 * Given the root of a binary tree and an integer targetSum, return all
 * root-to-leaf paths where the sum of the node values in the path equals
 * targetSum. Each path should be returned as a list of the node values, not
 * node references.
 *
 * A root-to-leaf path is a path starting from the root and ending at any leaf
 * node. A leaf is a node with no children.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
 * Output: [[5,4,11,2],[5,8,4,5]]
 * Explanation: There are two paths whose sum equals targetSum:
 * 5 + 4 + 11 + 2 = 22
 * 5 + 8 + 4 + 5 = 22
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,3], targetSum = 5
 * Output: []
 *
 *
 * Example 3:
 *
 *
 * Input: root = [1,2], targetSum = 0
 * Output: []
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 5000].
 * -1000 <= Node.val <= 1000
 * -1000 <= targetSum <= 1000
 *
 *
 */

#include "leetcode.h"
// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left),
 * right(right) {}
 * };
 */
class Solution {
public:
  vector<vector<int>> pathSum(TreeNode *root, int targetSum) {
    auto ret = vector<vector<int>>();
    if (!root) {
      return ret;
    }
    auto cur = vector<int>();
    cur.push_back(root->val);
    helper(root, targetSum - root->val, cur, ret);
    return ret;
  }
  void helper(TreeNode *root, int targetSum, vector<int> &cur,
              vector<vector<int>> &collect) {
    if (!root->left && !root->right) {
      if (targetSum == 0) {
        collect.push_back(cur);
      }
      return;
    }
    if (root->left) {
      cur.push_back(root->left->val);
      helper(root->left, targetSum - root->left->val, cur, collect);
      cur.pop_back();
    }
    if (root->right) {
      cur.push_back(root->right->val);
      helper(root->right, targetSum - root->right->val, cur, collect);
      cur.pop_back();
    }
  }
};
// @lc code=end

// [ 5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1 ]