/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    nums := []int{}

    var inOrder func(node *TreeNode)
    inOrder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inOrder(node.Left)
        nums = append(nums, node.Val)
        inOrder(node.Right)
    }
    inOrder(root)
    return nums[k-1]
}
