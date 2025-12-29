/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // if node exists - count it's length as 1 and pick max of left and right recursively
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

// concurrent solution, 
func maxDepth(root *TreeNode) int {
    input := make(chan int)
    go concurrentDFS(root, input)
    return <- input
}

func concurrentDFS(root *TreeNode, output chan int) {
    if root == nil {
        output <- 0
        return
    }
    input := make(chan int)
    go concurrentDFS(root.Left, input)
    go concurrentDFS(root.Right, input)
    output <- 1 + max(<-input, <-input) 
} 
