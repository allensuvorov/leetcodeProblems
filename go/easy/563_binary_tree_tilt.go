/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    // DFS sum of each sub tree
    totalTilt := 0
    var valueSum func(root *TreeNode) int
    valueSum = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        leftSum, rightSum := valueSum(root.Left), valueSum(root.Right)
        totalTilt += abs(leftSum - rightSum)
        return root.Val + leftSum + rightSum
    }
    valueSum(root)
    return totalTilt
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
