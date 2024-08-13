/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    ans := 100_000
    prev := -100_000
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        ans = min(ans, node.Val - prev)
        prev = node.Val
        dfs(node.Right)
    }
    dfs(root)
    return ans
}
