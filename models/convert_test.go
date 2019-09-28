package models

import "testing"

func TestCtoF(t *testing.T) {
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "CtoF",
			args: args{
				input: 0.0,
			},
			want: 32.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CtoF(tt.args.input); got != tt.want {
				t.Errorf("CtoF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFtoC(t *testing.T) {
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "FtoC", args: args{input: 32.0}, want: 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FtoC(tt.args.input); got != tt.want {
				t.Errorf("FtoC() = %v, want %v", got, tt.want)
			}
		})
	}
}
