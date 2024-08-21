/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    ans := [][]int{}
    q := []*TreeNode{root}
    direction := 1
    for len(q) > 0 {
        levLen := len(q)
        level := []int{}
        for range levLen {
            if direction == 1 {
                now := q[0]
                q = q[1:]
                level = append(level, now.Val)
                if now.Left != nil {
                    q = append(q, now.Left)
                }
                if now.Right != nil {
                    q = append(q, now.Right)
                }
            } else {
                top := len(q) - 1
                now := q[top]
                q = q[:top]
                level = append(level, now.Val)
                
                if now.Right != nil {
                    q = append([]*TreeNode{now.Right}, q...)
                }
                if now.Left != nil {
                    q = append([]*TreeNode{now.Left}, q...)
                }
            }
        }
        ans = append(ans, level)
        direction *= -1
    }
    return ans
}

// 9, 20 , 15, 7
