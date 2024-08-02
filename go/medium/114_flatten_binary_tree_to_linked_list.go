/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    reversePostOrderDFS(root)
}

func reversePostOrderDFS(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    tail := root
    if root.Right != nil && root.Left != nil {
        tail = reversePostOrderDFS(root.Right)
        reversePostOrderDFS(root.Left).Right = root.Right
        root.Right = root.Left
        root.Left = nil
    } else if root.Right != nil {
        tail = reversePostOrderDFS(root.Right)
    } else if root.Left != nil {
        tail = reversePostOrderDFS(root.Left)
        root.Right = root.Left
        root.Left = nil
    }
    return tail
}
