package main

import "testing"

func Test_loop1(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test loop 1",
			args: args{
				min: 1,
				max: 3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loop1(tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("loop1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loop2(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test loop 2",
			args: args{
				min: 7,
				max: 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loop2(tt.args.min, tt.args.max)
		})
	}
}

func Test_loop3(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test loop 3",
			args: args{
				min: 0,
				max: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loop3(tt.args.min, tt.args.max)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test for",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
