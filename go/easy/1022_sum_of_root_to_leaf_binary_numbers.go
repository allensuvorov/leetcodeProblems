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
    var dfs func(root *TreeNode, num int)
    dfs = func(root *TreeNode, num int) {
        if root == nil {
            return
        }
        num = num * 2 + root.Val
        
        if root.Left == nil && root.Right == nil {
            res += num
            return
        }
        dfs(root.Left, num)
        dfs(root.Right, num)
    }
    dfs(root, 0)
    return res
}
