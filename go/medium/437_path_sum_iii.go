/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	// dfs over dfs : time O (n^2)

	targetPathSumsCount := 0
	var countPathSums func(root *TreeNode, curSum int)
	countPathSums = func(root *TreeNode, curSum int) {
		if root == nil {
			return
		}
		curSum += root.Val
		if curSum == targetSum {
			targetPathSumsCount++
		}
		countPathSums(root.Left, curSum)
		countPathSums(root.Right, curSum)
	}
    
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		countPathSums(root, 0)
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	return targetPathSumsCount
}
