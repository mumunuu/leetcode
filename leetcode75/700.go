package leetcode75

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}

	if val < root.Val { // 더 작은곳으로 찾으러 가야함
		return searchBST(root.Left, val)
	}

	if val > root.Val { // 더 큰 곳으로 찾으러 가야함
		return searchBST(root.Right, val)
	}

	return root
}
