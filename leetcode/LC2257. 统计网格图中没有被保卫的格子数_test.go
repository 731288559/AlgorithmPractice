package leetcode

import "testing"

func Test_countUnguarded(t *testing.T) {
	type args struct {
		m      int
		n      int
		guards [][]int
		walls  [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				m:      4,
				n:      6,
				guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
				walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnguarded(tt.args.m, tt.args.n, tt.args.guards, tt.args.walls); got != tt.want {
				t.Errorf("countUnguarded() = %v, want %v", got, tt.want)
			}
		})
	}
}
