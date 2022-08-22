package hard

import "testing"

func Test_appealSum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				s: "abbca",
			},
			want: 28,
		},
		{
			name: "test2",
			args: args{
				s: "code",
			},
			want: 20,
		},
		{
			name: "test3",
			args: args{
				s: "vulong",
			},
			want: 56,
		},
		{
			name: "test4",
			args: args{
				s: "vulonggolang",
			},
			want: 279,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appealSum(tt.args.s); got != tt.want {
				t.Errorf("appealSum() = %v, want %v", got, tt.want)
			}
			if got := appealSum2(tt.args.s); got != tt.want {
				t.Errorf("appealSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOn2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appealSum("abbca")
		appealSum("code")
		appealSum("vulong")
		appealSum("vulonggolang")
	}
}

func Benchmark_appealSum_On2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appealSum("abbca")
		appealSum("code")
		appealSum("vulong")
		appealSum("vulonggolang")
	}
}

func Benchmark_appealSum_On(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appealSum2("abbca")
		appealSum2("code")
		appealSum2("vulong")
		appealSum2("vulonggolang")
	}
}