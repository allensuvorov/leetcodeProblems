/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    q := []*TreeNode{}
    if root != nil {
        q = append(q, root)
    }
    ans := [][]int{}
    for len(q) != 0 {
        rowLen := len(q)
        row := []int{}
        for range rowLen {
            now := q[0]
            q = q[1:]
            row = append(row, now.Val)
            if now.Left != nil {
                q = append(q, now.Left) 
            }
            if now.Right != nil { 
                q = append(q, now.Right) 
            }
        }
        ans = append(ans, row)
    }
    return ans
}
