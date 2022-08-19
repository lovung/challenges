package problems

import "testing"

func Test_kthSmallest(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	nums2 := []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				root: newBinaryTree([]*int{&nums[3], &nums[1], &nums[4], nil, &nums[2]}),
				k:    1,
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				root: newBinaryTree([]*int{&nums[5], &nums[3], &nums[6], &nums[2], &nums[4], nil, nil, &nums[1]}),
				k:    3,
			},
			want: 3,
		},
		{
			name: "test 3",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 1,
			},
			want: -10,
		},
		{
			name: "test 4",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 2,
			},
			want: -1,
		},
		{
			name: "test 5",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 3,
			},
			want: 3,
		},
		{
			name: "test 6",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 4,
			},
			want: 4,
		},
		{
			name: "test 7",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 5,
			},
			want: 5,
		},
		{
			name: "test 8",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 6,
			},
			want: 6,
		},
		{
			name: "test 9",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 7,
			},
			want: 7,
		},
		{
			name: "test 10",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 8,
			},
			want: 8,
		},
		{
			name: "test 11",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 9,
			},
			want: 9,
		},
		{
			name: "test 12",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 10,
			},
			want: 10,
		},
		{
			name: "test 13",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 11,
			},
			want: 11,
		},
		{
			name: "test 14",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 12,
			},
			want: 12,
		},
		{
			name: "test 15",
			args: args{
				root: newBinaryTree([]*int{&nums[10],
					&nums2[1], &nums[12],
					&nums2[10], &nums[7], &nums[11], &nums[13],
					nil, nil, &nums[5], &nums[8], nil, nil, nil, nil,
					nil, nil, nil, nil, &nums[4], &nums[6], nil, &nums[9], nil, nil, nil, nil, nil, nil, nil, nil,
					nil, nil, nil, nil, nil, nil, nil, nil, &nums[3],
				}),
				k: 13,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
