/*
 * @lc app=leetcode id=101 lang=cpp
 *
 * [101] Symmetric Tree
 *
 * https://leetcode.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (59.79%)
 * Likes:    16362
 * Dislikes: 428
 * Total Accepted:    2.6M
 * Total Submissions: 4.3M
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * Given the root of a binary tree, check whether it is a mirror of itself
 * (i.e., symmetric around its center).
 *
 *
 * Example 1:
 *
 *
 * Input: root = [1,2,2,3,4,4,3]
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: root = [1,2,2,null,3,null,3]
 * Output: false
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [1, 1000].
 * -100 <= Node.val <= 100
 *
 *
 *
 * Follow up: Could you solve it both recursively and iteratively?
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
//  the key is that we should change the problem to subtree mirror
class Solution {
public:
  bool isSymmetric(TreeNode *root) {
    if (!root) {
      return true;
    }
    return isMirror(root->left, root->right);
  }
  bool isMirror(TreeNode *tree1, TreeNode *tree2) {
    if (!tree1 && !tree2) {
      return true;
    }
    if (!tree1 || !tree2) {
      return false;
    }
    if (tree1->val != tree2->val) {
      return false;
    }
    return isMirror(tree1->left, tree2->right) &&
           isMirror(tree1->right, tree2->left);
  }
};
// @lc code=end
