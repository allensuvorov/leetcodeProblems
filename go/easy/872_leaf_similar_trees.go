/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leaves1 := []int{}
    dfs(root1, &leaves1)
    leaves2 := []int{}
    dfs(root2, &leaves2)


    if len(leaves1) != len(leaves2) {
        return false
    }

    for i := range leaves1 {
        if leaves1[i] != leaves2[i] {
            return false
        }
    }
    return true
}

func dfs(root *TreeNode, leaves *[]int){
    if root == nil {
        return
    }
    if root.Left == nil && root.Right == nil {
        *leaves = append(*leaves, root.Val)
    }
    dfs(root.Left, leaves)
    dfs(root.Right, leaves)
}
