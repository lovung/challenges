package apr2024

import "testing"

func Test_minRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			s:    "lee(t(c)ode",
			want: "lee(tc)ode",
		},
		{
			s:    "a)b(c)d",
			want: "ab(c)d",
		},
		{
			s:    "))((",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoveToMakeValid(tt.s); got != tt.want {
				t.Errorf("minRemoveToMakeValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
