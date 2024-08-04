/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var ( 
        sum int
        num int
        dfs func(root *TreeNode, num int)
    )
    
    dfs = func(root *TreeNode, num int) {
        if root == nil { return }
        num = num * 10 + root.Val
        if root.Left == nil && root.Right == nil {
            sum += num
        }
        dfs(root.Left, num)
        dfs(root.Right, num)
    }

    dfs(root, num)
    return sum
}
