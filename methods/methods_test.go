package main

import "testing"

func Test_rect_area(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test rect_area",
			fields: fields{
				width: 20,
				height: 10,
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
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Test rect_perim",
			fields: fields{
				width: 20,
				height: 10,
			},
			want: 60,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.perim(); got != tt.want {
				t.Errorf("rect.perim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test methods",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
