package leetcode

import "testing"

func Test_maxCollectedFruits(t *testing.T) {
	type args struct {
		fruits [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				fruits: [][]int{
					{1, 2, 3, 4},
					{5, 6, 8, 7},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: 100,
		},
		{
			name: "",
			args: args{
				fruits: [][]int{
					{1, 1},
					{1, 1},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCollectedFruits(tt.args.fruits); got != tt.want {
				t.Errorf("maxCollectedFruits() = %v, want %v", got, tt.want)
			}
		})
	}
}
