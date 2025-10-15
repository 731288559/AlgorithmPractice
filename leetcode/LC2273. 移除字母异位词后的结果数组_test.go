package leetcode

import (
	"reflect"
	"testing"
)

func Test_removeAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				words: []string{"abba", "baba", "bbaa", "cd", "cd"},
			},
			want: []string{"abba", "cd"},
		},
		{
			name: "",
			args: args{
				words: []string{"a", "b", "c", "d", "e"},
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeAnagrams(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
