package leetcode

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				board: [][]byte{
					{'X'},
				},
			},
		},
		{
			name: "",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
		})
	}
}
