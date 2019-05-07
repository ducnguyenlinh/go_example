package main

import "testing"

func Test_zeroval(t *testing.T) {
	type args struct {
		ival int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test zeroval",
			args: args{
				ival: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zeroval(tt.args.ival)
		})
	}
}

func Test_zeroptr(t *testing.T) {
	type args struct {
		iptr *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			zeroptr(tt.args.iptr)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test pointers",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
