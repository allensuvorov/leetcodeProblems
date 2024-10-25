/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    ans := 0
    prefixSumCount := map[int]int{0:1}
	var dfs func(root *TreeNode, sum int)
	dfs = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}
        sum += root.Val
        ans += prefixSumCount[sum - targetSum]
        prefixSumCount[sum]++
		dfs(root.Left, sum)
		dfs(root.Right, sum)
        prefixSumCount[sum]--
	}
	dfs(root, 0)
	return ans
}
