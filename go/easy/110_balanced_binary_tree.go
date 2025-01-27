/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return dfs(root) != -1
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    l, r := dfs(root.Left), dfs(root.Right)
    
    if l == -1 || r == -1 || abs(l, r) > 1 {
        return -1 
    }

    return 1+ max(l, r)
}

func abs(a, b int) int {
    if a - b > 0 {
        return a - b
    }
    return b - a
}
