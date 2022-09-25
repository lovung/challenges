package contest

import (
	"reflect"
	"testing"
)

func Test_sortPeople(t *testing.T) {
	type args struct {
		names   []string
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				names:   []string{"Mary", "John", "Emma"},
				heights: []int{180, 165, 170},
			},
			want: []string{"Mary", "Emma", "John"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPeople(tt.args.names, tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
