package weekly396

import "testing"

func Test_minimumOperationsToMakeKPeriodic(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				word: "leetcodeleet",
				k:    4,
			},
			want: 1,
		},
		{
			args: args{
				word: "leetcoleet",
				k:    2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumOperationsToMakeKPeriodic(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("minimumOperationsToMakeKPeriodic() = %v, want %v", got, tt.want)
			}
		})
	}
}
