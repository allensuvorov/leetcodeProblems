/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    node := &TreeNode{Val:preorder[0]}
    m := slices.Index(inorder, preorder[0])
    node.Left = buildTree(preorder[1:m+1], inorder[:m])
    node.Right = buildTree(preorder[m+1:], inorder[m+1:])
    return node
}
