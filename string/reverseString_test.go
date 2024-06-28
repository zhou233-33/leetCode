package string

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "reverseString",
			args: args{
				s: []byte{'h', 'e', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'e', 'h'},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := ReverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ReverseString() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestReverseStringFrontK(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := ReverseStringFrontK(tt.args.s, tt.args.k); got != tt.want {
					t.Errorf("ReverseStringFrontK() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
