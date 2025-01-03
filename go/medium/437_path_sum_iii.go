/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
    prefSumSet := map[int]int{0:1}
    res := 0
    var dfs func(root *TreeNode, sum int)
    dfs = func(root *TreeNode, sum int) {
        if root == nil {
            return
        }
        sum += root.Val
        res += prefSumSet[sum - targetSum]
        prefSumSet[sum]++
        dfs(root.Left, sum)
        dfs(root.Right, sum)
        prefSumSet[sum]--
    }
    dfs(root, 0)
    return res
}
