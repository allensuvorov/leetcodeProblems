/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
    ans := 0
    maxSum := math.MinInt
    level := 1
    q := []*TreeNode{root}
    
    for len(q) > 0 {
        levelSize := len(q)
        sum := 0
        for range levelSize {
            now := q[0]
            q = q[1:]
            
            if now.Left != nil {
                q = append(q, now.Left)
            }
            if now.Right != nil {
                q = append(q, now.Right)
            }
            sum += now.Val
        }
        if sum > maxSum {
            maxSum = sum
            ans = level
        }
        level++
    }
    return ans
}
