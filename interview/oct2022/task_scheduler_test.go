package oct2022

import (
	"reflect"
	"testing"
)

func Test_executeTask(t *testing.T) {
	type args struct {
		tasks []*task
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				tasks: []*task{
					{"T1", 10, 10, 50},
					{"T2", 0, 15, 10},
					{"T3", 20, 5, 100},
					{"T4", 15, 5, 10},
					{"T5", 0, 15, 5},
				},
			},
			want: []string{"T2", "T1", "T3", "T4", "T5"},
		},
		{
			args: args{
				tasks: []*task{
					{"T1", 10, 10, 50},
					{"T2", 5, 15, 10},
					{"T3", 20, 5, 100},
					{"T4", 15, 5, 10},
					{"T5", 5, 15, 5},
				},
			},
			want: []string{"T2", "T3", "T1", "T4", "T5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeTask(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("executeTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
