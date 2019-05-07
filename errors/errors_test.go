package main

import "testing"

func Test_f1(t *testing.T) {
	type args struct {
		arg int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test f1 error",
			args: args{
				arg: 42,
			},
			want: -1,
			wantErr: true,
		},
		{
			name: "Test f1 no error",
			args: args{
				arg: 7,
			},
			want: 10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f1(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("f1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("f1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_argError_Error(t *testing.T) {
	type fields struct {
		arg  int
		prob string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test argError error",
			fields: fields{
				arg: 42,
				prob: "can't work with it",
			},
			want: "42 - can't work with it",
		},
		{
			name: "Test argError no error",
			fields: fields{
				arg: 7,
				prob: "",
			},
			want: "7 - ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &argError{
				arg:  tt.fields.arg,
				prob: tt.fields.prob,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("argError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_f2(t *testing.T) {
	type args struct {
		arg int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test f2 error",
			args: args{
				arg: 42,
			},
			want: -1,
			wantErr: true,
		},
		{
			name: "Test f2 no error",
			args: args{
				arg: 7,
			},
			want: 10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f2(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("f2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("f2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test errors",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
