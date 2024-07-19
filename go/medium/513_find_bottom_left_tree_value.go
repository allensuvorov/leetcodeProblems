/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
    q := []*TreeNode{root}
    for len(q) > 0 {
        curRowJobsNum := len(q)
        isLastRow := true
        for i := range curRowJobsNum {
            now := q[i]
            if now.Left != nil {
                isLastRow = false
                q = append(q, now.Left)

            }

            if now.Right != nil {
                isLastRow = false
                q = append(q, now.Right)
            }
        }
        if isLastRow {
            return q[0].Val
        }
        q = q[curRowJobsNum:]

    }
    return 0
}
