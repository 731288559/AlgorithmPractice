package leetcode

import (
	"testing"
)

func Test_numOfMinutes(t *testing.T) {
	type args struct {
		n          int
		headID     int
		manager    []int
		informTime []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TEST-1",
			args: args{
				n:          1,
				headID:     0,
				manager:    []int{-1},
				informTime: []int{0},
			},
			want: 0,
		},
		{
			name: "TEST-2",
			args: args{
				n:          6,
				headID:     2,
				manager:    []int{2, 2, -1, 2, 2, 2},
				informTime: []int{0, 0, 1, 0, 0, 0},
			},
			want: 1,
		},
		{
			name: "TEST-3",
			args: args{
				n:          9,
				headID:     2,
				manager:    []int{2, 2, -1, 2, 2, 2, 0, 1, 7},
				informTime: []int{5, 2, 1, 0, 0, 0, 0, 2, 0},
			},
			want: 6,
		},
		{
			name: "TEST-4",
			args: args{
				n:          9,
				headID:     2,
				manager:    []int{2, 2, -1, 2, 2, 2, 0, 1, 7},
				informTime: []int{3, 2, 1, 0, 0, 0, 0, 2, 0},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfMinutes(tt.args.n, tt.args.headID, tt.args.manager, tt.args.informTime); got != tt.want {
				t.Errorf("numOfMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
