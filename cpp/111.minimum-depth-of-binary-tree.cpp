/*
 * @lc app=leetcode id=111 lang=cpp
 *
 * [111] Minimum Depth of Binary Tree
 */
#include<leetcode-helper.h>
// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
   
public:
    Solution(){
    min = 100000;
    }
    int minDepth(TreeNode* root) {
        minDepthHelper(root, 0);
        return min;
    }
    void minDepthHelper(TreeNode* root,int cur){
        if (root==nullptr){
            if (min>cur){
                min = cur;
            }
            return;
        }
        minDepthHelper(root->left,cur + 1);
        minDepthHelper(root->right, cur + 1);
    }

private:
    int min;
};
// @lc code=end

