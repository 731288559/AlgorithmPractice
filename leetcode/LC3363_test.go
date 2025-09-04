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
			name: "case 1",
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
			name: "case 2",
			args: args{
				fruits: [][]int{
					{1, 1},
					{1, 1},
				},
			},
			want: 4,
		},
		{
			name: "case 3",
			args: args{
				fruits: [][]int{
					{8, 5, 0, 17, 15},
					{6, 0, 15, 6, 0},
					{7, 19, 16, 8, 18},
					{11, 3, 2, 12, 13},
					{17, 15, 15, 4, 6},
				},
			},
			want: 145,
		},
		{
			name: "case 4",
			args: args{
				fruits: [][]int{
					{16, 3, 11, 14, 14},
					{3, 0, 10, 13, 14},
					{7, 18, 8, 7, 18},
					{7, 8, 5, 7, 5},
					{0, 14, 8, 1, 0},
				},
			},
			want: 105,
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
