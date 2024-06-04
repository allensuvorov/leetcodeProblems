/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
    var newRoot *TreeNode
    var dfs func(node *TreeNode) *TreeNode
    dfs = func(node *TreeNode) *TreeNode {
        if node == nil {
            return nil
        }

        if node.Left != nil && newRoot == nil {
            if node.Left.Left == nil {
                newRoot = node.Left
            }
        }

        left := dfs(node.Left)
        if left != nil {
            left.Right = node
        }

        right := dfs(node.Right)
        if right != nil {
            return right
        }
        return node
    }    

    dfs(root)
    if newRoot == nil {
        return root
    }
    return newRoot
}
