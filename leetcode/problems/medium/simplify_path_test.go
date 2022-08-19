package medium

import "testing"

func Test_simplifyPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "test2",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "test3",
			args: args{
				path: "/a/../../b/../c//.//",
			},
			want: "/c",
		},
		{
			name: "test4",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "test5",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPath(tt.args.path); got != tt.want {
				t.Errorf("simplifyPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
