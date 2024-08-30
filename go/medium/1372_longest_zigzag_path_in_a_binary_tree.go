/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode, zig bool) int
    dfs = func(node *TreeNode, zig bool) int {
        if node == nil {
            return 0
        }
        left := dfs(node.Left, false)
        right := dfs(node.Right, true)
        if zig {
            res = max(res, right, left + 1)
            return left + 1
        }
        res = max(res, left, right + 1)
        return right + 1
    }
    dfs(root.Left, false)
    dfs(root.Right, true)
    return res 
}
