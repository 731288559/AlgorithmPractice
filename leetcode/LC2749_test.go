package leetcode

import "testing"

func Test_makeTheIntegerZero(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				num1: 3,
				num2: -2,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				num1: 5,
				num2: 7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeTheIntegerZero(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("makeTheIntegerZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
