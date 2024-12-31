/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var dfs func(node *TreeNode, leaves *[]int) *[]int
    
    dfs = func(node *TreeNode, leaves *[]int) *[]int {
        if node == nil {
            return nil
        }
        left := dfs(node.Left, leaves)
        right := dfs(node.Right, leaves)
        if left == nil && right == nil {
            *leaves = append(*leaves, node.Val)
        }
        return leaves
    }

    leftLeaves := dfs(root1, &[]int{})
    rightLeaves := dfs(root2, &[]int{})
    fmt.Println(*leftLeaves, *rightLeaves)

    return slices.Compare(*leftLeaves, *rightLeaves) == 0
}

