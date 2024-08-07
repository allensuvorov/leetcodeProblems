/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return dfs(root.Left, root.Right)
}

func dfs(root1, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil {
        return true
    }
    if root1 == nil || root2 == nil || root1.Val != root2.Val {
        return false
    }
    dfs(root1.Left, root2.Right)
    return dfs(root1.Left, root2.Right) && dfs(root1.Right, root2.Left)
}
