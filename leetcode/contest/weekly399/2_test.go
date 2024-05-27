package weekly399

import "testing"

func Test_compressedString(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				word: "abcde",
			},
			want: "1a1b1c1d1e",
		},
		{
			args: args{
				word: "aaaaaaaaaaaaaabb",
			},
			want: "9a5a2b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressedString(tt.args.word); got != tt.want {
				t.Errorf("compressedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
