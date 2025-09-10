package leetcode

import "testing"

func Test_peopleAwareOfSecret(t *testing.T) {
	type args struct {
		n      int
		delay  int
		forget int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n:      6,
				delay:  2,
				forget: 4,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				n:      4,
				delay:  1,
				forget: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peopleAwareOfSecret(tt.args.n, tt.args.delay, tt.args.forget); got != tt.want {
				t.Errorf("peopleAwareOfSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}
