package leetcode

import (
	"reflect"
	"testing"
)

func Test_getSneakyNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums: []int{7, 1, 5, 4, 3, 4, 6, 0, 9, 5, 8, 2},
			},
			want: []int{5, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSneakyNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSneakyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
