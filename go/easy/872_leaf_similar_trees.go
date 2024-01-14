/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaves1 := make([]int, 0)
    leaves2 := make([]int, 0)

    dfs(root1, &leaves1)
    dfs(root2, &leaves2)

    return slices.Compare(leaves1, leaves2) == 0
}

func dfs(node *TreeNode, leaves *[]int) {
    if node == nil {
        return
    }
    if node.Left == nil && node.Right == nil {
        *leaves = append(*leaves, node.Val)
    } else {
        dfs(node.Left, leaves)
        dfs(node.Right, leaves)
    }
}
