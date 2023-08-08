package leetcode

import "testing"

func Test_maxAbsoluteSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, -3, 2, 3, -4},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				nums: []int{2, -5, 1, -4, 3, -2},
			},
			want: 8,
		},
		{
			name: "3",
			args: args{
				nums: []int{1, -3, 10},
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAbsoluteSum(tt.args.nums); got != tt.want {
				t.Errorf("maxAbsoluteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
