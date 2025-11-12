package leetcode

import "testing"

func Test_maximumProfit(t *testing.T) {
	type args struct {
		prices []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				prices: []int{1, 7, 9, 8, 2},
				k:      2,
			},
			want: 14,
		},
		{
			name: "",
			args: args{
				prices: []int{12, 16, 19, 19, 8, 1, 19, 13, 9},
				k:      3,
			},
			want: 36,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProfit(tt.args.prices, tt.args.k); got != tt.want {
				t.Errorf("maximumProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
