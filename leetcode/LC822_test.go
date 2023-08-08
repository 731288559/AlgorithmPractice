package leetcode

import "testing"

func Test_flipgame(t *testing.T) {
	type args struct {
		fronts []int
		backs  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				fronts: []int{1, 2, 4, 4, 7},
				backs:  []int{1, 3, 4, 1, 3},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				fronts: []int{1},
				backs:  []int{1},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				fronts: []int{1, 1},
				backs:  []int{1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipgame(tt.args.fronts, tt.args.backs); got != tt.want {
				t.Errorf("flipgame() = %v, want %v", got, tt.want)
			}
		})
	}
}
