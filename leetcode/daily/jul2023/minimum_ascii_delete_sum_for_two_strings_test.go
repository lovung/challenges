package jul2023

import (
	"testing"
)

func Test_minimumDeleteSum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
		{
			args: args{
				s1: "bbccacacaab",
				s2: "aabccb",
			},
			want: 878,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumDeleteSum2(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "sea",
				s2: "eat",
			},
			want: 231,
		},
		{
			args: args{
				s1: "delete",
				s2: "leet",
			},
			want: 403,
		},
		{
			args: args{
				s1: "bbccacacaab",
				s2: "aabccb",
			},
			want: 878,
		},
		{
			args: args{
				s1: "igijekdtywibepwonjbwykkqmrgmtybwhwjiqudxmnniskqjfbkpcxukrablqmwjndlhblxflgehddrvwfacarwkcpmcfqnajqfxyqwiugztocqzuikamtvmbjrypfqvzqiwooewpzcpwhdejmuahqtukistxgfafrymoaodtluaexucnndlnpeszdfsvfofdylcicrrevjggasrgdhwdgjwcchyanodmzmuqeupnpnsmdkcfszznklqjhjqaboikughrnxxggbfyjriuvdsusvmhiaszicfa",
				s2: "ikhuivqorirphlzqgcruwirpewbjgrjtugwpnkbrdfufjsmgzzjespzdcdjcoioaqybciofdzbdieegetnogoibbwfielwungehetanktjqjrddkrnsxvdmehaeyrpzxrxkhlepdgpwhgpnaatkzbxbnopecfkxoekcdntjyrmmvppcxcgquhomcsltiqzqzmkloomvfayxhawlyqxnsbyskjtzxiyrsaobbnjpgzmetpqvscyycutdkpjpzfokvi",
			},
			want: 41731,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeleteSum2(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumDeleteSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIfCanBeTrimmedString(t *testing.T) {
	type args struct {
		s  string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s:  "let",
				s2: "leet",
			},
			want: true,
		},
		{
			args: args{
				s:  "acb",
				s2: "accbxxxd",
			},
			want: true,
		},
		{
			args: args{
				s:  "acb",
				s2: "accdxxx",
			},
			want: false,
		},
		{
			args: args{
				s:  "lete",
				s2: "leet",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfCanBeTrimmedString(tt.args.s, tt.args.s2); got != tt.want {
				t.Errorf("checkIfCanBeTrimmedString(%s, %s) = %v, want %v", tt.args.s, tt.args.s2, got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequenceASCIISum(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s1: "let",
				s2: "leet",
			},
			want: sumASCII("let"),
		},
		{
			args: args{
				s1: "acb",
				s2: "accbxxxd",
			},
			want: sumASCII("acb"),
		},
		{
			args: args{
				s1: "acb",
				s2: "accdxxx",
			},
			want: sumASCII("ac"),
		},
		{
			args: args{
				s1: "lete",
				s2: "leet",
			},
			want: sumASCII("let"),
		},
		{
			args: args{
				s1: "bbccacacaab",
				s2: "aabccb",
			},
			want: sumASCII("bccb"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceMaxASCIISum(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSubsequenceASCIISum() = %v, want %v", got, tt.want)
			}
		})
	}
}
