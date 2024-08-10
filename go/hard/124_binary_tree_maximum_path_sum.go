/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
    ans := root.Val
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := dfs(root.Left)
        right := dfs(root.Right)
        curMax := max(root.Val, root.Val + max(left, right))
        ans = max(ans, curMax, left + right + root.Val)
        return curMax
    }
    dfs(root)
    return ans
}
