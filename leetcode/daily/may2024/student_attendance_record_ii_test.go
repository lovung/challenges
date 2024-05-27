package may2024

import "testing"

func Test_checkRecord(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			args: 4,
			want: 43,
		},
		{
			args: 5,
			want: 94,
		},
		{
			args: 10101,
			want: 183236316,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRecord(tt.args); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			_ = checkRecord_wrong(tt.args)
		})
	}
}
