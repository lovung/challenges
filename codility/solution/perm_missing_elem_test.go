package solution

import "testing"

func TestPermMissingElem(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				A: []int{3, 1, 2, 5},
			},
			want: 4,
		},
		{
			name: "test 2",
			args: args{
				A: []int{3, 4, 2, 5},
			},
			want: 1,
		},
		{
			name: "test 3",
			args: args{
				A: []int{3, 4, 2, 1},
			},
			want: 5,
		},
		{
			name: "test 4",
			args: args{
				A: []int{2},
			},
			want: 1,
		},
		{
			name: "test 5",
			args: args{
				A: []int{},
			},
			want: 1,
		},
		{
			name: "test 6",
			args: args{
				A: []int{1},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PermMissingElem(tt.args.A); got != tt.want {
				t.Errorf("PermMissingElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
