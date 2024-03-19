package leetcode

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				num: 2736,
			},
			want: 7236,
		},
		{
			name: "",
			args: args{
				num: 7236,
			},
			want: 7632,
		},
		{
			name: "",
			args: args{
				num: 9973,
			},
			want: 9973,
		},
		{
			name: "",
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				num: 115,
			},
			want: 511,
		},
		{
			name: "",
			args: args{
				num: 7222216,
			},
			want: 7622212,
		},
		{
			name: "",
			args: args{
				num: 1993,
			},
			want: 9913,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
