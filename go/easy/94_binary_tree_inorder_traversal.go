/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    left := inorderTraversal(root.Left)
    right := inorderTraversal(root.Right)
    return append(append(left, root.Val), right...)
}
