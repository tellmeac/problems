package lcode

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			args: args{
				s: "cbaebabacd",
				p: "abc",
			},
			wantResult: []int{0, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("findAnagrams() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
