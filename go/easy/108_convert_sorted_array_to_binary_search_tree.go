/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {

    var dfs func(l, r int) *TreeNode
    dfs = func(l, r int) *TreeNode {
        if l < 0 || r >= len(nums) || l > r {
            return nil
        }
        node := new(TreeNode)
        mid := l + (r - l) / 2 
        node.Val = nums[mid]
        
        node.Left = dfs(l, mid - 1)
        node.Right = dfs(mid + 1, r)
        return node
    }

    return dfs(0, len(nums) - 1)
}
