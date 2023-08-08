package leetcode

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				time: []int{30, 20, 150, 100, 40},
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				time: []int{60, 60, 60},
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				time: []int{30, 20, 30, 40, 40, 20},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
