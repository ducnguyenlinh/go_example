package main

import "testing"

func Test_vals(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			name: "Test vals",
			want: 3,
			want1: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := vals()
			if got != tt.want {
				t.Errorf("vals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("vals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test multiple return values",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
