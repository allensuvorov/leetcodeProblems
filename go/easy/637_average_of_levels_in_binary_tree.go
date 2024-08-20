/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    res := []float64{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        rowLen := len(q)
        sum := 0
        for range rowLen {
            now := q[0]
            q = q[1:]

            sum += now.Val

            if now.Left != nil {
                q = append(q, now.Left)
            }

            if now.Right != nil {
                q = append(q, now.Right)
            }
        }
        res = append(res, float64(sum) / float64(rowLen))
    }
    return res
}
