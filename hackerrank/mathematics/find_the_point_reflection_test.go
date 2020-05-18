package mathematics

import (
	"reflect"
	"testing"
)

func Test_findPoint(t *testing.T) {
	type args struct {
		px int32
		py int32
		qx int32
		qy int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "test 1",
			args: args{
				px: 0,
				py: 0,
				qx: 1,
				qy: 1,
			},
			want: []int32{2, 2},
		},
		{
			name: "test 2",
			args: args{
				px: 1,
				py: 1,
				qx: 2,
				qy: 2,
			},
			want: []int32{3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPoint(tt.args.px, tt.args.py, tt.args.qx, tt.args.qy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
