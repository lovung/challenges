package contest

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums:    []int{4, 5, 2, 1},
				queries: []int{3, 10, 21},
			},
			want: []int{2, 3, 4},
		},
		{
			name: "test2",
			args: args{
				nums:    []int{4, 5, 2, 3},
				queries: []int{1},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
