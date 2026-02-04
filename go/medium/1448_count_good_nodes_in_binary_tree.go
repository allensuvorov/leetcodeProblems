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

func dfs(root *TreeNode, good int) int {
    if root == nil {
        return 0
    }
    res := 0
    if root.Val >= good {
        res += 1
        good = root.Val
    }
    
    l := dfs(root.Left, good)
    r := dfs(root.Right, good)

    return res + l + r
}

