package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test range_over_channels",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
