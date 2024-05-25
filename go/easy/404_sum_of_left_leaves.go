/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return root.Val
    }
    left := sumOfLeftLeaves(root.Left)
    right := sumOfLeftLeaves(root.Right)
    return left + right
}
