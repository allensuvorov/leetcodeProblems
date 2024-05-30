/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {
        return nil
    }

    node := new(TreeNode)

    var root1Left, root1Right *TreeNode

    if root1 != nil {
        node.Val += root1.Val
        root1Left, root1Right = root1.Left, root1.Right
    }
    
    var root2Left, root2Right *TreeNode

    if root2 != nil {
        node.Val += root2.Val
        root2Left, root2Right = root2.Left, root2.Right
    }
    
    node.Left = mergeTrees(root1Left, root2Left )
    node.Right = mergeTrees(root1Right, root2Right)

    return node
}
