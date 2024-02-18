/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root.Val > p.Val && root.Val > q.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    if root.Val < p.Val && root.Val < q.Val{
        return lowestCommonAncestor(root.Right, p, q)
    }
    return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
        switch {
        case root.Val > p.Val && root.Val > q.Val:
            root = root.Left
        case root.Val < p.Val && root.Val < q.Val:
            root = root.Right
        default: 
            return root
        }
    }
    return nil
}
