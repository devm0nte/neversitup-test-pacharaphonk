package main

import (
	"reflect"
	"testing"
)

func TestCountSmilyFace(t *testing.T) {
	type args struct {
		text []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "given [] result should be 0",
			args: args{text: []string{}},
			want: 0,
		},
		{
			name: "given [':)', ';(', ';}', ':-D'] result should be 2",
			args: args{text: []string{":)", ";(", ";}", ":-D"}},
			want: 2,
		},
		{
			name: "given [';D', ':-(', ':-)', ';~)'] result should be 3",
			args: args{text: []string{";D", ":-(", ":-)", ";~)"}},
			want: 3,
		},
		{
			name: "given [';]', ':[', ';*', ':$', ';-D'] result should be 1",
			args: args{text: []string{";]", ":[", ";*", ":$", ";-D"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSmilyFace(tt.args.text); got != tt.want {
				t.Errorf("CountSmilyFace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTheFace(t *testing.T) {
	type args struct {
		face string
	}
	tests := []struct {
		name  string
		args  args
		want  theFace
		want1 bool
	}{
		{
			name:  "given ':)' result should be theFace Struct and true",
			args:  args{face: ":)"},
			want:  theFace{Eye: ":", Nose: "", Mouth: ")"},
			want1: true,
		},
		{
			name:  "given :( result should be theFace Struct and true",
			args:  args{face: ":("},
			want:  theFace{Eye: ":", Nose: "", Mouth: "("},
			want1: true,
		},
		{
			name:  "given ;-} result should be theFace Struct and true",
			args:  args{face: ";-}"},
			want:  theFace{Eye: ";", Nose: "-", Mouth: "}"},
			want1: true,
		},
		{
			name:  "given * result should be theFace Struct and true",
			args:  args{face: "*"},
			want:  theFace{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetTheFace(tt.args.face)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTheFace() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetTheFace() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsValidSmilyFace(t *testing.T) {
	type args struct {
		validFace theFace
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "given ;-} result should be false",
			args: args{validFace: theFace{Eye: ";", Nose: "-", Mouth: "}"}},
			want: false,
		},
		{
			name: "given :> result should be false",
			args: args{validFace: theFace{Eye: ":", Nose: "", Mouth: ">"}},
			want: false,
		},
		{
			name: "given :) result should be true",
			args: args{validFace: theFace{Eye: ":", Nose: "", Mouth: ")"}},
			want: true,
		},
		{
			name: "given ;-D result should be true",
			args: args{validFace: theFace{Eye: ";", Nose: "-", Mouth: "D"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSmilyFace(tt.args.validFace); got != tt.want {
				t.Errorf("IsValidSmilyFace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	testSlice := []string{"a", "b", "c", "d", "999"}
	type args struct {
		slice []string
		t     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "given testSlice and t is 'b' result should be true",
			args: args{testSlice, "b"},
			want: true,
		},
		{
			name: "given testSlice and t is 'ab' result should be false",
			args: args{testSlice, "ab"},
			want: false,
		},
		{
			name: "given testSlice and t is '999' result should be true",
			args: args{testSlice, "999"},
			want: true,
		},
		{
			name: "given testSlice and t is '9' result should be false",
			args: args{testSlice, "9"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.slice, tt.args.t); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
