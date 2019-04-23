package main

import (
	"testing"
)

func Test_switch1(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test switch1",
			args: args{
				i: 3,
			},
			want: "three",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := switch1(tt.args.i); got != tt.want {
				t.Errorf("switch1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check_weekday(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test check_weekday",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check_weekday()
		})
	}
}

func Test_check_hour(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test check_hour",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check_hour()
		})
	}
}

func Test_check_type(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test check_type",
			args: args{
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check_type(tt.args.i)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test switch",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
