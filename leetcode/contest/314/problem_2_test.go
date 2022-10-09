package contest

import (
	"reflect"
	"testing"
)

func Test_findArray(t *testing.T) {
	type args struct {
		pref []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				pref: []int{5, 2, 0, 3, 1},
			},
			want: []int{5, 7, 2, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findArray(tt.args.pref); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
