/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
     if root == nil {
        return 0
     }
     maxSubTreeDepth := 0
     for _, v := range root.Children {
        maxSubTreeDepth = max(maxSubTreeDepth, maxDepth(v))
     }
     return maxSubTreeDepth + 1
}
