package leetcode

import "testing"

func Test_sumOfPower(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{2, 1, 4},
			},
			want: 141,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPower2(tt.args.nums); got != tt.want {
				t.Errorf("sumOfPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
