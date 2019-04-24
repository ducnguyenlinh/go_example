package main

import "testing"

func Test_plus(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test plus",
			args: args{
				a: 4,
				b: 6,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plus(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("plus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_plusPlus(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test plusPlus",
			args: args{
				a: 2,
				b: 4,
				c: 6,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusPlus(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("plusPlus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test functions",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
