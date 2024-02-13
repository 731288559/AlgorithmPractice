package leetcode

import (
	"reflect"
	"testing"
)

func Test_replaceValueInTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 10},
					},
					Right: &TreeNode{
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val:   0,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Right: &TreeNode{Val: 11},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceValueInTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceValueInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
