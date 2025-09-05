package leetcode

import "testing"

func Test_numberOfPairs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1",
			args: args{
				points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			},
			want: 0,
		},
		{
			name: "case 2",
			args: args{
				points: [][]int{{6, 3}, {4, 4}, {2, 6}},
			},
			want: 2,
		},
		{
			name: "case 3",
			args: args{
				points: [][]int{{3, 1}, {1, 3}, {1, 1}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfPairs(tt.args.points); got != tt.want {
				t.Errorf("numberOfPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
