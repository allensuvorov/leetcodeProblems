/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if left, right := getDepth(root.Left), getDepth(root.Right); left == right {
        return 1 << left + countNodes(root.Right)
    } else {
        return 1 << right + countNodes(root.Left)
    }
    return 0
}

func getDepth(node *TreeNode) int {
    if node == nil {
        return 0
    }
    return getDepth(node.Left) + 1
}
