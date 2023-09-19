package leetcode

import "testing"

func Test_rob3(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: &TreeNode{Val: 1},
					},
				},
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  &TreeNode{Val: 1},
						Right: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: &TreeNode{Val: 1},
					},
				},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob3(tt.args.root); got != tt.want {
				t.Errorf("rob3() = %v, want %v", got, tt.want)
			}
		})
	}
}
