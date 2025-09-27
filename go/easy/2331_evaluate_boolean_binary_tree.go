/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    if root.Val < 2 {
        return root.Val == 1
    }
    
    left := evaluateTree(root.Left)
    right := evaluateTree(root.Right)

    if root.Val == 2 {
        return left || right
    }

    return left && right    
}
