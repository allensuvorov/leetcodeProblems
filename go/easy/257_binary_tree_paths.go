/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    paths := dfs(root, "")
    return paths
}

func dfs(root *TreeNode, path string) []string {
    if root == nil {
        return nil
    }
    if len(path) > 0 {
        path += "->"
    } 
    path += strconv.Itoa(root.Val)
    if root.Left == nil && root.Right == nil {
        return []string{path} 
    }
    left := dfs(root.Left, path)
    right := dfs(root.Right, path)
    return append(left, right...)
} 
