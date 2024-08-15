/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    prev := math.MinInt
    ans := true
    var inOrder func(node *TreeNode)
    inOrder = func(node *TreeNode) {
        if node == nil {
            return 
        }
        inOrder(node.Left)
        if prev >= node.Val {
            ans = false
            return
        }
        prev = node.Val
        inOrder(node.Right)
    }
    inOrder(root)
    return ans
}
