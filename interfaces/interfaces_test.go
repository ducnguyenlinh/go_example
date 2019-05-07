package main

import (
	"testing"
)

func Test_rect_area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test rect_area",
			fields: fields{
				width: 10,
				height: 20,
			},
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.area(); got != tt.want {
				t.Errorf("rect.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rect_perim(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test rect_perim",
			fields: fields{
				width: 10,
				height: 20,
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.perim(); got != tt.want {
				t.Errorf("rect.perim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_circle_area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test circle_area",
			fields: fields{
				radius: 3,
			},
			want: 28.274333882308138,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := circle{
				radius: tt.fields.radius,
			}
			if got := c.area(); got != tt.want {
				t.Errorf("circle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_circle_perim(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Test circle_perim",
			fields: fields{
				radius: 3,
			},
			want: 18.84955592153876,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := circle{
				radius: tt.fields.radius,
			}
			if got := c.perim(); got != tt.want {
				t.Errorf("circle.perim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_measure(t *testing.T) {
	type args struct {
		g geometry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			measure(tt.args.g)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test interfaces",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
