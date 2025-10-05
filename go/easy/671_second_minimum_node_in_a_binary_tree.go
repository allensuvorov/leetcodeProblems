/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    min1 := root.Val
    min2 := math.MaxInt

    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root != nil {
            if root.Val > min1 {
                min2 = min(min2, root.Val)
            } else {
                dfs(root.Left)
                dfs(root.Right)
            }
        }
    }

    dfs(root)
    
    if min2 < math.MaxInt {
        return min2
    }

    return -1
}
