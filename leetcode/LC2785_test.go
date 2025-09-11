package leetcode

import "testing"

func Test_sortVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				s: "lEetcOde",
			},
			want: "lEOtcede",
		},
		{
			name: "",
			args: args{
				s: "lYmpH",
			},
			want: "lYmpH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortVowels(tt.args.s); got != tt.want {
				t.Errorf("sortVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
