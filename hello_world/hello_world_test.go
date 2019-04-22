package main

import "testing"

func Test_hello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test func hello",
			want: "Hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello(); got != tt.want {
				t.Errorf("hello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
