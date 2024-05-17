/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {         
    if root == nil { return root }
    if removeLeafNodes(root.Left, target) == nil { root.Left = nil }
    if removeLeafNodes(root.Right, target) == nil { root.Right = nil }
    if root.Left == nil && root.Right == nil && root.Val == target { return nil }
    return root
}
