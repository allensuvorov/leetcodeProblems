/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode, path int)
    dfs = func(root *TreeNode, path int) {
        if root == nil {
            return
        }

        path = path << 1 + root.Val

        if root.Left == nil && root.Right == nil {
            res = res + path
        }

        dfs(root.Left, path)
        dfs(root.Right, path)
    }

    dfs(root, 0)
    return res
}
