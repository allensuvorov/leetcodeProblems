/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode { 
    
    var DFS func(root *TreeNode) bool
    DFS = func(root *TreeNode) bool {
        if root == nil { return false }
        if DFS(root.Left) { root.Left = nil }
        if DFS(root.Right) { root.Right = nil }
        if root.Left == nil && root.Right == nil && root.Val == target { return true }
        return false
    }
    if DFS(root) { return nil }
    return root
}
