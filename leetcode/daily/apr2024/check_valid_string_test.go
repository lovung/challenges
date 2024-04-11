package apr2024

import "testing"

func Test_checkValidString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			s:    ")",
			want: false,
		},
		{
			s:    "((*)(*))((*",
			want: false,
		},
		{
			s:    "(())((())()()(*)(*()(())())())()()((()())((()))(*",
			want: false,
		},
		{
			s:    "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())",
			want: false,
		},
		{
			s:    "(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString1(tt.s); got != tt.want {
				t.Errorf("checkValidString1() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
