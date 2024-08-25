/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    left := postorderTraversal(root.Left)
    right := postorderTraversal(root.Right)
    return append(left, append(right, root.Val)...)
}
