/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ans := []int{}
    var dfs func(root *TreeNode) []int
    dfs = func (root *TreeNode) []int {
        if root == nil {
            return nil
        }
        dfs(root.Left)
        ans = append(ans, root.Val)
        dfs(root.Right)
        return ans
    }
    dfs(root)
    return ans
}
