package mathematics

import "testing"

func Test_summingSeries(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "test 1",
			args: args{
				n: 5351871996120528,
			},
			want: 578351320,
		},
		{
			name: "test 2",
			args: args{
				n: 2,
			},
			want: 4,
		},
		{
			name: "test 3",
			args: args{
				n: 2248813659738258,
			},
			want: 404664464,
		},
		{
			name: "test 3",
			args: args{
				n: 2248813659738258,
			},
			want: 404664464,
		},
		{
			name: "test 3",
			args: args{
				n: 2248813659738258,
			},
			want: 404664464,
		},
		{
			name: "test 4",
			args: args{
				n: 2494359640703601,
			},
			want: 20752136,
		},
		{
			name: "test 5",
			args: args{
				n: 6044763399160734,
			},
			want: 999516029,
		},
		{
			name: "test 6",
			args: args{
				n: 3271269997212342,
			},
			want: 743537718,
		},
		{
			name: "test 7",
			args: args{
				n: 4276346434761561,
			},
			want: 323730244,
		},
		{
			name: "test 8",
			args: args{
				n: 2372239019637533,
			},
			want: 174995256,
		},
		{
			name: "test 9",
			args: args{
				n: 5624204919070546,
			},
			want: 593331567,
		},
		{
			name: "test 10",
			args: args{
				n: 9493965694520825,
			},
			want: 136582381,
		},
		{
			name: "test 11",
			args: args{
				n: 8629828692375133,
			},
			want: 305527433,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summingSeries(tt.args.n); got != tt.want {
				t.Errorf("summingSeries() = %v, want %v", got, tt.want)
			}
		})
	}
}
