package feature

import "testing"

func TestWaitGroupTest2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitGroupTest2()
		})
	}
}

func TestWaitGroupTest3(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitGroupTest3()
		})
	}
}
