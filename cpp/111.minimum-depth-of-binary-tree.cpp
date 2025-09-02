/*
 * @lc app=leetcode id=111 lang=cpp
 *
 * [111] Minimum Depth of Binary Tree
 *
 * https://leetcode.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (51.23%)
 * Likes:    7649
 * Dislikes: 1352
 * Total Accepted:    1.5M
 * Total Submissions: 2.9M
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, find its minimum depth.
 *
 * The minimum depth is the number of nodes along the shortest path from the
 * root node down to the nearest leaf node.
 *
 * Note: A leaf is a node with no children.
 *
 *
 * Example 1:
 *
 *
 * Input: root = [3,9,20,null,null,15,7]
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: root = [2,null,3,null,4,null,5,null,6]
 * Output: 5
 *
 *
 *
 * Constraints:
 *
 *
 * The number of nodes in the tree is in the range [0, 10^5].
 * -1000 <= Node.val <= 1000
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

// 和MaxDepth 最大的区别在于必须是叶子结点
//        1
//      2
//    3
//  4
//  这样的树的最小深度应该是 4

class Solution {
public:
  int minDepth(TreeNode *root) {
    if (!root) {
      return 0;
    }
    // 叶子结点的高度是 0
    int height = 0;
    if (!root->left && !root->right) {
      height = 1;
    } else if (!root->left) {
      height = minDepth(root->right) + 1;
    } else if (!root->right) {
      height = minDepth(root->left) + 1;
    } else {
      auto left = minDepth(root->left);
      auto right = minDepth(root->right);
      height = min(left, right) + 1;
    }
    return height;
  }
};
// @lc code=end
// 2 3
//   3
// 9   20
//   15    7

// [ 2, null, 3, null, 4, null, 5, null, 6 ]
// 2
//   3
//     4
//       5
//         6

TEST(minDepth, case1) {
  TreeNode *node2 = new TreeNode(2);
  TreeNode *node3 = new TreeNode(3);
  TreeNode *node4 = new TreeNode(4);
  TreeNode *node5 = new TreeNode(5);
  TreeNode *node6 = new TreeNode(6);
  node2->right = node3;
  node3->right = node4;
  node4->right = node5;
  node5->right = node6;
  Solution().minDepth(node2);
}