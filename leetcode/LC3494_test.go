package leetcode

import "testing"

func Test_minTime(t *testing.T) {
	type args struct {
		skill []int
		mana  []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				skill: []int{1, 5, 2, 4},
				mana:  []int{5, 1, 4, 2},
			},
			want: 110,
		},
		{
			name: "",
			args: args{
				skill: []int{1, 1, 1},
				mana:  []int{1, 1, 1},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				skill: []int{1, 2, 3, 4},
				mana:  []int{1, 2},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTime(tt.args.skill, tt.args.mana); got != tt.want {
				t.Errorf("minTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
