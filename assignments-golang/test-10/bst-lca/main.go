package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{}
	left := &TreeNode{}
	right := &TreeNode{}
	root.Left = left
	root.Right = right
}
func lca(root, v1, v2) {
	if root == nil {

	}
}
