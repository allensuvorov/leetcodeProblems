/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorderTraversal(root *TreeNode) []int {
    answer := make([]int, 0, 100)
    p := &answer
    dfs(root, p)
    return answer
}

func dfs(node *TreeNode, p *[]int) {
    if node == nil {
        return
    }
    *p = append(*p, node.Val)
    dfs(node.Left, p)
    dfs(node.Right, p)
}
