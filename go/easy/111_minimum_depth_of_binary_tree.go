// DFS solution
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
    depth := 0
    if root == nil {
        return 0
    }
    var DFS func (root *TreeNode)
    DFS = func(root *TreeNode) {
        if root == nil {
            return
        }
        depth++
        if root.Left == nil && root.Right == nil {
            md = min(md, depth)
        }
        DFS(root.Left)
        DFS(root.Right)
        depth--
        return
    }
    DFS(root)
    return md
}

// BFS, counting distance till first leaf
