/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    if root == p || root == q {
        return root
    }

    l := lowestCommonAncestor(root.Left, p, q)
    r := lowestCommonAncestor(root.Right, p, q)

    switch {
    case l != nil && r != nil: // lca case
        return root
    case l != nil:
        return l
    }
    return r
}
