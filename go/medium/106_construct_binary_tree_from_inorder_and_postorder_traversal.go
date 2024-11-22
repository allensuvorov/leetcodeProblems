/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 || len(postorder) == 0 {
        return nil
    }
    n := len(postorder)
    node := &TreeNode {Val: postorder[n-1]}
    m := slices.Index(inorder, postorder[n-1])
    node.Left = buildTree(inorder[:m], postorder[:m])
    node.Right = buildTree(inorder[m+1:], postorder[m:n-1])
    return node
}
