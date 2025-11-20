/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // if node exists - count it's length as 1 and pick max of left and right recursively
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
