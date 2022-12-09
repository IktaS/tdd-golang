package tdd_golang_test

import (
	"testing"

	tdd_golang "github.com/supriyadi687/tdd-golang"
)

func TestShouldReturn0GivenEmptyString(t *testing.T) {
	res, _ := tdd_golang.Add("")
	if res != 0 {
		t.Errorf("expected to be 0 got %d", res)
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test should return 0 given empty string",
			args: args{
				input: "",
			},
			want: 0,
		},
		{
			name: "test should return 1 given one number, 1",
			args: args{
				input: "1",
			},
			want: 1,
		},
		{
			name: "test should return 2 given one number, 2",
			args: args{
				input: "2",
			},
			want: 2,
		},
		{
			name: "test should return sum = 3 given two number, 1, 2",
			args: args{
				input: "1,2",
			},
			want: 3,
		},
		{
			name: "test should return sum = 17 given five number, 1, 2,3,5,6",
			args: args{
				input: "1,2,3,5,6",
			},
			want: 17,
		},
		{
			name: "test should return sum = 14 given three number, 3,5,6",
			args: args{
				input: "3,5,6",
			},
			want: 14,
		},
		{
			name: "test should return sum = 14 given three number with newline, 3,5,6",
			args: args{
				input: "3,\n5,6",
			},
			want: 14,
		},
		{
			name: "test should return sum = 14 given three number with newline,no comma, 3,5,6",
			args: args{
				input: "3\n5,6",
			},
			want: 14,
		},
		{
			name: "test should return sum = 8 given two number, with delimiter ; | 3;5",
			args: args{
				input: "//;\n3;5",
			},
			want: 8,
		},
		{
			name: "test should return sum = 14 given two number, with delimiter ;, with newline | 3;5",
			args: args{
				input: "//;\n3\n5",
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tdd_golang.Add(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
