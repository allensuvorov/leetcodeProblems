/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    res := 0
    level := 0
    maxLevelSum := math.MinInt
    q := []*TreeNode{root}
    for len(q) > 0 {
        levelSize := len(q)
        levelSum := 0
        level++
        for range levelSize {
            now := q[0]
            q = q[1:]
            levelSum += now.Val
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
        }
        if levelSum > maxLevelSum {
            maxLevelSum = levelSum
            res = level
        }
    }
    return res
}
