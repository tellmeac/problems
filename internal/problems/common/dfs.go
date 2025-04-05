package common

// TreeNode for binary tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode) []int {
	var route []int
	f := func(val int) {
		route = append(route, val)
	}

	var move func(*TreeNode)
	move = func(x *TreeNode) {
		if x == nil {
			return
		}

		move(x.Left)
		f(x.Value)
		move(x.Right)
	}

	move(root)

	return route
}

func postorder(root *TreeNode) []int {
	var route []int
	f := func(val int) {
		route = append(route, val)
	}

	var move func(*TreeNode)
	move = func(x *TreeNode) {
		if x == nil {
			return
		}

		move(x.Left)
		move(x.Right)
		f(x.Value)
	}

	move(root)

	return route
}

func preorder(root *TreeNode) []int {
	var route []int
	f := func(val int) {
		route = append(route, val)
	}

	var move func(*TreeNode)
	move = func(x *TreeNode) {
		if x == nil {
			return
		}

		f(x.Value)
		move(x.Left)
		move(x.Right)
	}

	move(root)

	return route
}
