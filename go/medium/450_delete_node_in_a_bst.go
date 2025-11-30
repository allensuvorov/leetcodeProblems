/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	}
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	}

	if key == root.Val {
		// key == root.Val case
		// one child or no children cases
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// we have two children - find next bigger
		nextBigger := root.Right
		for nextBigger.Left != nil {
			nextBigger = nextBigger.Left
		}
		root.Val = nextBigger.Val
		root.Right = deleteNode(root.Right, nextBigger.Val)
	}

	return root
}
