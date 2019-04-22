package main

import "testing"

func Test_check_even_odd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test check_even_odd",
			args: args{
				n: 4,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check_even_odd(tt.args.n); got != tt.want {
				t.Errorf("check_even_odd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check_divisible(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test check_divisibale",
			args: args{
				a: 6,
				b: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := check_divisible(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("check_divisible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_check_multiple_digits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test check_multiple_digits",
			args: args{
				num: 11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check_multiple_digits(tt.args.num)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test if else",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
