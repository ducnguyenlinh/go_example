package main

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test sum",
			args: args{
				nums: []int{2, 4, 6, 8},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.nums); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_range_maps(t *testing.T) {
	type args struct {
		kvs map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test range_maps",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			range_maps(tt.args.kvs)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test range",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
