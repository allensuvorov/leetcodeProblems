/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    maxSum := 0
    
    var dfs func (node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }

        left := dfs(node.Left)
        right := dfs(node.Right)
        
        // update maxSum
        maxSum = max(maxSum, left + right)
        
        // return maxBranch
        return max(left, right) + 1
    }
    dfs(root)
    return maxSum
}


