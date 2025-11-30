/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    maxSum := math.MinInt
    res := 0 // result
    q := []*TreeNode{root}
    x := 1 // level
    for len(q) > 0 {
        size := len(q)
        curSum := 0
        for range size {
            now := q[0]
            q = q[1:]
            curSum += now.Val
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
        }
        if curSum > maxSum {
            maxSum = curSum
            res = x
        }
        x++
    }
    return res
}
