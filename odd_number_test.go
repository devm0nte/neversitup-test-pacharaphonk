package main

import (
	"testing"
)

func TestIsOdd(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "given 0 result should be false",
			args: args{number: 0},
			want: false,
		},
		{
			name: "given 1 result should be true",
			args: args{number: 1},
			want: true,
		},
		{
			name: "given -101 result should be true",
			args: args{number: -101},
			want: true,
		},
		{
			name: "given 102 result should be false",
			args: args{number: 102},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOdd(tt.args.number); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOddNumber(t *testing.T) {
	type args struct {
		text []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given [7] result should be 7",
			args: args{text: []int{7}},
			want: 7,
		},
		{
			name: "given [0] result should be 0",
			args: args{text: []int{0}},
			want: 0,
		},
		{
			name: "given [1,1,2] result should be 2",
			args: args{text: []int{1, 1, 2}},
			want: 2,
		},
		{
			name: "given [0,1,0,1,0] result should be 0",
			args: args{text: []int{0, 1, 0, 1, 0}},
			want: 0,
		},
		{
			name: "given [1,2,2,3,3,3,4,3,3,3,2,2,1] result should be 4",
			args: args{text: []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOddNumber(tt.args.text); got != tt.want {
				t.Errorf("FindOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
