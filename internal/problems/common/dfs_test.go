package common

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDFS(t *testing.T) {
	tree := &TreeNode{
		Value: 1,
		Left:  &TreeNode{Value: 2},
		Right: &TreeNode{Value: 3},
	}

	for _, tt := range []struct {
		name    string
		dfsFunc func(*TreeNode) []int
		root    *TreeNode
		want    []int
	}{
		{
			name:    "inorder",
			root:    tree,
			dfsFunc: inorder,
			want:    []int{2, 1, 3},
		},
		{
			name:    "preorder",
			root:    tree,
			dfsFunc: preorder,
			want:    []int{1, 2, 3},
		},
		{
			name:    "postorder",
			root:    tree,
			dfsFunc: postorder,
			want:    []int{2, 3, 1},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.dfsFunc(tt.root)

			require.Equal(t, tt.want, got)
		})
	}
}
