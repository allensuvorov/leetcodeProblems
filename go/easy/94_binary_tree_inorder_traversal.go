/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    l := inorderTraversal(root.Left)
    r := inorderTraversal(root.Right)
    return append(
        append(l, root.Val),
        r...,
    )
}

// concurrent processing of left and right at each node.
func inorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var wg sync.WaitGroup
    var l, r []int
    wg.Add(1)
    go func(){
        l = inorderTraversal(root.Left)
        wg.Done()
    }()
    wg.Add(1)
    go func(){
        r = inorderTraversal(root.Right)
        wg.Done()
    }()
    wg.Wait()
    return append(append(l, root.Val), r...)
}
