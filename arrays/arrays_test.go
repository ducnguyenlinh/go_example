package main

import "testing"

func Test_one_dimensionaly_array(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test one_dimensional",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			one_dimensionaly_array()
		})
	}
}

func Test_multi_dimensional_array(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test multi_dimensional",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			multi_dimensional_array()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test arrays",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
