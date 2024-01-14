/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return dfs(root, root.Val)
}

func dfs(node *TreeNode, max int) int {
    if node == nil {
        return 0
    }
    good := 0
    if max <= node.Val {
        good = 1
        max = node.Val
    }
    return good + dfs(node.Left, max) + dfs(node.Right, max)
}
