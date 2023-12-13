package main

import (
	"reflect"
	"testing"
)

func TestManipulate(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "given '' result should be empty slice",
			args: args{text: ""},
			want: []string{},
		},
		{
			name: "given 'a' result should be ['a']",
			args: args{text: "a"},
			want: []string{"a"},
		},
		{
			name: "given 'ab' result should be ['ab','ba']",
			args: args{text: "ab"},
			want: []string{"ab", "ba"},
		},
		{
			name: "given 'abc' result should be ['abc','acb','bac','bca','cab','cba']",
			args: args{text: "abc"},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name: "given 'aabb' result should be ['aabb','abab','abba','baab','baba','bbaa']",
			args: args{text: "abc"},
			want: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Manipulate(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ManipulateInputString() = %v, want %v", got, tt.want)
			}
		})
	}
}
