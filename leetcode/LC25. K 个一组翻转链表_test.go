package leetcode

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					}},
				k: 2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			tmp := got
			i := 0
			for tmp != nil {
				assert.Equal(t, tt.want[i], tmp.Val)
				i++
				tmp = tmp.Next
			}
			assert.Equal(t, i, len(tt.want))
		})
	}
}
