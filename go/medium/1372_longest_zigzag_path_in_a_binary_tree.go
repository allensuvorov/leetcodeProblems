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
    var dfs func(root *TreeNode, l, r int)
    dfs = func(root *TreeNode, l, r int) {
        if root == nil {
            return
        }
        res = max(res, l, r)
        dfs(root.Left, r+1, 0)
        dfs(root.Right, 0, l+1)
    }

    dfs(root, 0, 0)
    return res
}
