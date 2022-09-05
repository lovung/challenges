package binarytree

// func Test_verticalTraversal(t *testing.T) {
// 	type args struct {
// 		root *trees.TreeNode[int]
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want [][]int
// 	}{
// 		{
// 			args: args{
// 				root: trees.Slice2TreeNode[int](
// 					[]*int{
// 						pointer.Of(3), pointer.Of(9), pointer.Of(20),
// 						nil, nil, pointer.Of(15), pointer.Of(7),
// 					},
// 				),
// 			},
// 			want: [][]int{
// 				{9}, {3, 15}, {20}, {7},
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := verticalTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("verticalTraversal() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
