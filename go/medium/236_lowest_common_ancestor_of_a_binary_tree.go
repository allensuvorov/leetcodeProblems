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
    if l != nil && r != nil {
        return root
    }
    if l != nil {
        return l
    }
    return r
}

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

    l := lowestCommonAncestor(root.Left, p, q)
    r := lowestCommonAncestor(root.Right, p, q)

    switch { 
    case l == p && r == q, l == q && r == p:
        return root
    case root == p && (l == q || r == q):
        return root
    case root == q && (l == p || r == p):
        return root
    case root == p, l == p, r == p: 
        return p
    case root == q, l == q, r == q:
        return q
    case l != nil:
        return l
    case r != nil:
        return r
    }
    return nil
}
