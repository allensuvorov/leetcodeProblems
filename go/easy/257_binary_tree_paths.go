/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return []string{strconv.Itoa(root.Val)} 
    }
    left := binaryTreePaths(root.Left)
    right := binaryTreePaths(root.Right)
    paths := append(left, right...)
    for i := range paths {
        paths[i] = strconv.Itoa(root.Val) + "->" + paths[i]
    }
    return paths
}
