package feature

import "testing"

func TestChannelTest2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChannelTest2()
		})
	}
}

func TestChannelTest1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ChannelTest1()
		})
	}
}
