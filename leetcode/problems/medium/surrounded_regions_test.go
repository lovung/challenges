package medium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "test1",
			args: args{
				board: [][]byte{
					{x, x, x, x},
					{x, o, o, x},
					{x, x, o, x},
					{x, o, x, x},
				},
			},
			want: [][]byte{
				{x, x, x, x},
				{x, x, x, x},
				{x, x, x, x},
				{x, o, x, x},
			},
		},
		{
			name: "test2",
			args: args{
				board: [][]byte{
					{o, o, o},
					{o, o, o},
					{o, o, o},
				},
			},
			want: [][]byte{
				{o, o, o},
				{o, o, o},
				{o, o, o},
			},
		},
		{
			name: "test3",
			args: args{
				board: [][]byte{
					{x, x, x},
					{x, o, x},
					{x, x, x},
				},
			},
			want: [][]byte{
				{x, x, x},
				{x, x, x},
				{x, x, x},
			},
		},
		{
			name: "test4",
			args: args{
				board: [][]byte{
					{x},
				},
			},
			want: [][]byte{
				{x},
			},
		},
		{
			name: "test5",
			args: args{
				board: [][]byte{
					{o},
				},
			},
			want: [][]byte{
				{o},
			},
		},
		{
			name: "test6",
			args: args{
				board: [][]byte{},
			},
			want: [][]byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			assert.Equal(t, tt.want, tt.args.board)
		})
	}
}

func Test_solveUnionFind(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "test1",
			args: args{
				board: [][]byte{
					{x, x, x, x},
					{x, o, o, x},
					{x, x, o, x},
					{x, o, x, x},
				},
			},
			want: [][]byte{
				{x, x, x, x},
				{x, x, x, x},
				{x, x, x, x},
				{x, o, x, x},
			},
		},
		{
			name: "test2",
			args: args{
				board: [][]byte{
					{o, o, o},
					{o, o, o},
					{o, o, o},
				},
			},
			want: [][]byte{
				{o, o, o},
				{o, o, o},
				{o, o, o},
			},
		},
		{
			name: "test3",
			args: args{
				board: [][]byte{
					{x, x, x},
					{x, o, x},
					{x, x, x},
				},
			},
			want: [][]byte{
				{x, x, x},
				{x, x, x},
				{x, x, x},
			},
		},
		{
			name: "test4",
			args: args{
				board: [][]byte{
					{x},
				},
			},
			want: [][]byte{
				{x},
			},
		},
		{
			name: "test5",
			args: args{
				board: [][]byte{
					{o},
				},
			},
			want: [][]byte{
				{o},
			},
		},
		{
			name: "test6",
			args: args{
				board: [][]byte{},
			},
			want: [][]byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveWithUnionFind(tt.args.board)
			assert.Equal(t, tt.want, tt.args.board)
		})
	}
}
