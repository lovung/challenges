package weekly396

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		word string
		want bool
	}{
		{
			word: "234Adas",
			want: true,
		},
		{
			word: "b3",
			want: false,
		},
		{
			word: "Db3",
			want: false,
		},
		{
			word: "a3$e",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.word); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
