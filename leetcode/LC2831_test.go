package leetcode

import "testing"

func Test_longestEqualSubarray(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 3, 2, 3, 1, 3},
				k:    3,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2, 2, 1, 1},
				k:    2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestEqualSubarray(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestEqualSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
