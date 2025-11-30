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

    if root == p || root == q { // i am once of targets and could be LCA
        return root
    }

    l := lowestCommonAncestor(root.Left, p, q)
    r := lowestCommonAncestor(root.Right, p, q)

    if l != nil && r != nil { // i am the LCA
        return root
    }

    if l != nil {
        return l
    }

    if r != nil {
        return r
    }
     
    return nil
}
