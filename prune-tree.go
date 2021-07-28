/*
Given the root of a binary tree, return the same tree where every subtree (of the given tree) not containing a 1 has been removed.

A subtree of a node node is node plus every node that is a descendant of node.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    
    if root.Left != nil {
        root.Left = pruneTree(root.Left)    
    }
    if root.Right != nil {
        root.Right = pruneTree(root.Right)
    }
    if root.Left == nil && root.Right == nil {
        if root.Val == 0 {
            return nil
        }
        return root
    }
    return root
}
