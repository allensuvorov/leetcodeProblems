/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    md := 100000
    if root == nil {
        return 0
    }
    return DFS(root, md, 1)
}

func DFS(root *TreeNode, md, d int) int {
    if root == nil {
        return md
    }
    if root.Left == nil && root.Right == nil {
        return min(md, d)
    }
    return min(DFS(root.Left, md, d + 1), DFS(root.Right, md, d + 1))
}
