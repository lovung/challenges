package contest

import (
	"reflect"
	"testing"
)

func Test_buildMatrix(t *testing.T) {
	type args struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				k:             3,
				rowConditions: [][]int{{1, 2}, {3, 2}},
				colConditions: [][]int{{2, 1}, {3, 2}},
			},
			want: [][]int{
				{0, 0, 1},
				{3, 0, 0},
				{0, 2, 0},
			},
		},
		{
			name: "test2",
			args: args{
				k:             3,
				rowConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}},
				colConditions: [][]int{{2, 1}},
			},
			want: [][]int{},
		},
		{
			name: "test3",
			args: args{
				k:             3,
				rowConditions: [][]int{{2, 1}},
				colConditions: [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}},
			},
			want: [][]int{},
		},
		{
			name: "test4",
			args: args{
				k:             8,
				rowConditions: [][]int{{1, 2}, {7, 3}, {4, 3}, {5, 8}, {7, 8}, {8, 2}, {5, 8}, {3, 2}, {1, 3}, {7, 6}, {4, 3}, {7, 4}, {4, 8}, {7, 3}, {7, 5}},
				colConditions: [][]int{{5, 7}, {2, 7}, {4, 3}, {6, 7}, {4, 3}, {2, 3}, {6, 2}},
			},
			want: [][]int{
				{1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 7, 0},
				{0, 0, 0, 6, 0, 0, 0, 0},
				{0, 4, 0, 0, 0, 0, 0, 0},
				{0, 0, 5, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 3},
				{0, 0, 0, 0, 8, 0, 0, 0},
				{0, 0, 0, 0, 0, 2, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildMatrix(tt.args.k, tt.args.rowConditions, tt.args.colConditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toposortKahnAlgo(t *testing.T) {
	type args struct {
		k   int
		con [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				k:   8,
				con: [][]int{{1, 2}, {7, 3}, {4, 3}, {5, 8}, {7, 8}, {8, 2}, {5, 8}, {3, 2}, {1, 3}, {7, 6}, {4, 3}, {7, 4}, {4, 8}, {7, 3}, {7, 5}},
			},
			want: []int{1, 7, 6, 4, 5, 3, 8, 2},
		},
		{
			name: "test2",
			args: args{
				k:   3,
				con: [][]int{{1, 2}, {2, 3}, {3, 2}},
			},
			want: []int{1},
		},
		{
			name: "test3",
			args: args{
				k: 298,
				con: [][]int{{102, 221}, {19, 76}, {173, 106}, {61, 192}, {257, 135}, {15, 162}, {27, 244}, {15, 277}, {33, 91}, {134, 114}, {264, 72}, {213, 298}, {247, 236}, {255, 281}, {246, 277}, {246, 165},
					{152, 20}, {179, 1}, {194, 298}, {271, 253}, {181, 2}, {230, 70}, {103, 206}, {169, 60}, {76, 94}, {102, 84}, {141, 7}, {138, 153}, {290, 46}, {286, 138}, {119, 33}, {226, 31}, {141, 11},
					{183, 267}, {202, 29}, {282, 167}, {40, 220}, {247, 89}, {145, 84}, {273, 198}, {286, 121}, {282, 295}, {115, 274}, {247, 60}, {194, 292}, {146, 211}, {121, 261}, {156, 150}, {51, 44},
					{86, 142}, {42, 224}, {95, 260}, {192, 135}, {173, 167}, {12, 242}, {235, 127}, {146, 169}, {238, 286}, {295, 25}, {228, 226}, {251, 283}, {177, 144}, {113, 163}, {224, 58}, {199, 225},
					{76, 17}, {101, 298}, {148, 231}, {242, 73}, {159, 90}, {268, 168}, {60, 82}, {180, 131}, {176, 210}, {141, 292}, {30, 216}, {126, 108}, {127, 157}, {244, 243}, {227, 241}, {251, 80},
					{112, 23}, {282, 8}, {182, 277}, {238, 53}, {44, 138}, {14, 87}, {138, 170}, {203, 36}, {169, 95}, {182, 229}, {88, 69}, {8, 48}, {31, 23}, {156, 103}, {198, 74}, {173, 41}, {42, 285},
					{204, 295}, {41, 195}, {195, 223}, {112, 206}, {173, 87}, {166, 126}, {233, 149}, {267, 220}, {287, 175}, {217, 261}, {145, 99}, {181, 207}, {63, 54}, {132, 278}, {166, 220}, {183, 149},
					{51, 195}, {195, 101}, {288, 97}, {292, 20}, {173, 34}, {185, 298}, {110, 35}, {295, 144}, {20, 275}, {42, 125}, {161, 169}, {48, 235}, {150, 5}, {124, 250}, {14, 255}, {147, 12},
					{271, 229}, {28, 35}, {266, 11}, {242, 250}, {32, 262}, {217, 165}, {237, 95}, {43, 275}, {246, 161}, {86, 292}, {193, 16}, {242, 216}, {293, 84}, {298, 190}, {182, 129}, {251, 113},
					{263, 267}, {33, 217}, {37, 78}, {166, 35}, {249, 174}, {158, 222}, {267, 112}, {94, 211}, {137, 51}, {143, 46}, {231, 36}, {177, 174}, {154, 18}, {89, 231}, {41, 79}, {202, 265}, {232, 170},
					{156, 149}, {277, 159}, {83, 75}, {231, 219}, {296, 159}, {173, 62}, {238, 46}, {176, 164}, {180, 106}, {148, 289}, {131, 90}, {231, 267}, {248, 43}, {40, 225}, {213, 196}, {89, 127},
					{71, 159}, {174, 175}, {127, 82}, {74, 82}, {207, 22}, {2, 279}, {120, 262}, {119, 88}, {46, 257}, {133, 265}, {259, 285}, {215, 58}, {3, 37}, {63, 171}, {39, 104}, {21, 18}, {61, 128},
					{131, 173}, {241, 26}, {88, 266}, {201, 87}, {7, 201}, {128, 92}, {91, 245}, {246, 180}, {236, 243}, {132, 210}, {63, 244}, {286, 14}, {284, 49}, {132, 230}, {86, 183}, {147, 291},
					{148, 71}, {264, 239}, {166, 169}, {286, 115}, {249, 140}, {21, 14}, {33, 72}, {176, 249}, {195, 16}, {271, 58}, {53, 256}, {263, 252}, {56, 103}, {152, 199}, {214, 190}, {225, 125},
					{183, 266}, {152, 19}, {161, 293}, {290, 164}, {106, 160}, {37, 190}, {266, 154}, {95, 74}, {55, 137}, {245, 195}, {40, 202}, {218, 137}, {41, 139}, {93, 47}, {44, 99}, {258, 225},
					{83, 98}, {279, 140}, {287, 149}, {76, 146}, {51, 256}, {55, 180}, {152, 237}, {198, 107}, {259, 165}, {71, 280}, {205, 47}, {55, 295}, {143, 171}, {252, 72}, {235, 51}, {122, 76},
					{228, 227}, {273, 76}, {237, 109}, {66, 295}, {92, 116}, {227, 3}, {13, 120}, {85, 245}, {108, 291}, {57, 197}, {260, 4}, {17, 3}, {113, 233}, {296, 54}, {181, 162}, {202, 59},
					{235, 244}, {296, 162}, {278, 276}, {239, 54}, {211, 3}, {9, 133}, {200, 86}, {235, 11}, {200, 88}, {285, 182}, {110, 225}, {173, 269}, {189, 51}, {76, 227}, {74, 140}, {166, 170},
					{258, 233}, {9, 194}, {180, 45}, {40, 168}, {292, 73}, {182, 290}, {60, 4}, {99, 225}, {189, 287}, {284, 142}, {44, 118}, {57, 26}, {66, 142}, {195, 143}, {184, 27}, {121, 9}, {146, 290},
					{145, 49}, {37, 97}, {188, 269}, {98, 179}, {69, 196}, {191, 190}, {8, 232}, {108, 291}, {194, 12}, {217, 43}, {29, 124}, {2, 55}, {251, 17}, {255, 240}, {121, 254}, {146, 148},
					{230, 159}, {235, 200}, {56, 212}, {15, 218}, {48, 244}, {215, 231}, {269, 87}, {279, 77}, {238, 110}, {161, 79}, {259, 160}, {241, 31}, {158, 283}, {279, 236}, {12, 203}, {285, 277},
					{3, 111}, {100, 216}, {53, 4}, {218, 206}, {198, 282}, {284, 237}, {106, 22}, {229, 178}, {104, 195}, {55, 72}, {281, 219}, {27, 272}, {169, 78}, {9, 240}, {76, 188}, {182, 231},
					{246, 239}, {184, 240}, {28, 130}, {182, 252}, {94, 31}, {24, 162}, {243, 87}, {110, 46}, {126, 206}, {103, 160}, {218, 27}, {289, 295}, {169, 10}, {13, 133}, {259, 209}, {152, 239},
					{109, 96}, {263, 144}, {233, 78}, {15, 282}, {235, 113}, {77, 111}, {68, 206}, {148, 252}, {218, 297}, {254, 159}, {88, 135}, {198, 281}, {45, 20}, {39, 254}, {54, 196}, {265, 135}, {131, 59}, {207, 38}, {180, 147}, {142, 120}, {104, 243}, {14, 179}, {40, 134}, {99, 49}, {8, 202},
					{154, 68}, {236, 158}, {73, 108}, {69, 288}, {129, 95}, {116, 26}, {147, 122}, {242, 28}, {150, 57}, {104, 46}, {200, 173}, {55, 89}, {254, 201}, {85, 266}, {226, 3}, {273, 244}, {285, 199},
					{228, 68}, {48, 186}, {85, 110}, {132, 155}, {169, 165}, {204, 297}, {234, 118}, {247, 219}, {112, 108}, {158, 105}, {63, 56}, {104, 248}, {84, 230}, {174, 120}, {17, 263}, {54, 136},
					{296, 276}, {188, 203}, {296, 25}, {164, 32}, {263, 20}, {249, 239}, {14, 230}, {115, 144}, {5, 118}, {179, 197}, {210, 153}, {110, 59}, {281, 206}, {221, 20}, {66, 229}, {17, 289},
					{271, 152}, {65, 99}, {43, 154}, {247, 7}, {209, 199}, {241, 259}, {166, 63}, {152, 19}, {53, 256}, {238, 283}, {249, 97}, {198, 158}, {165, 72}, {203, 215}, {227, 111}, {9, 23},
					{79, 36}, {42, 71}, {106, 160}, {237, 7}, {213, 265}, {29, 53}, {256, 101}, {124, 73}, {237, 166}, {282, 101}, {287, 70}, {261, 267}, {259, 85}, {103, 130}, {255, 270}, {26, 288},
					{13, 254}, {246, 247}, {28, 256}, {23, 240}, {279, 120}, {86, 25}, {11, 208}, {193, 23}, {284, 54}, {77, 251}, {1, 192}, {57, 52}, {277, 133}, {282, 206}, {182, 220}, {234, 140},
					{172, 54}, {232, 82}, {106, 199}, {175, 165}, {5, 190}, {224, 23}, {129, 141}, {181, 82}, {204, 195}, {107, 66}, {215, 231}, {166, 21}, {91, 266}, {107, 212}, {152, 193}, {169, 129},
					{131, 243}, {278, 250}, {88, 295}, {215, 52}, {163, 32}, {72, 190}, {186, 225}, {236, 68}, {234, 31}, {251, 184}, {156, 11}, {164, 212}, {297, 36}, {176, 107}, {83, 285}, {138, 1},
					{112, 171}, {226, 175}, {76, 272}, {41, 136}, {148, 52}, {221, 17}, {42, 126}, {219, 143}, {204, 256}, {174, 298}, {107, 185}, {238, 286}, {286, 121}, {277, 212}, {281, 30}, {95, 197},
					{66, 119}, {75, 114}, {106, 135}, {268, 243}, {259, 284}, {264, 196}, {259, 127}, {205, 262}, {242, 83}, {113, 86}, {86, 206}, {179, 36}, {173, 108}, {21, 231}, {33, 199}, {10, 24},
					{120, 72}, {287, 239}, {228, 71}, {30, 16}, {168, 193}, {222, 220}, {92, 159}, {245, 169}, {268, 225}, {39, 241}, {8, 127}, {13, 157}, {182, 257}, {60, 178}, {132, 92}, {33, 215}, {111, 47},
					{66, 10}, {222, 81}, {55, 29}, {2, 107}, {51, 277}, {69, 274}, {156, 230}, {30, 185}, {89, 84}, {15, 186}, {39, 57}, {261, 40}, {132, 114}, {228, 160}, {10, 216}, {144, 190}, {131, 102},
					{267, 10}, {50, 176}, {98, 46}, {74, 108}, {143, 157}, {166, 47}, {55, 70}, {267, 140}, {194, 59}, {77, 185}, {2, 31}, {237, 215}, {226, 187}, {117, 272}, {81, 153}, {266, 184}, {287, 221},
					{246, 242}, {173, 35}, {146, 103}, {228, 247}, {13, 39}, {48, 134}, {169, 205}, {19, 279}, {85, 236}, {248, 204}, {40, 172}, {88, 81}, {218, 162}, {100, 280}, {176, 102}, {151, 211},
					{38, 275}, {177, 151}, {2, 213}, {251, 277}, {88, 169}, {168, 4}, {127, 44}, {258, 54}, {221, 132}, {211, 17}, {88, 70}, {251, 180}, {293, 75}, {254, 23}, {51, 45}, {5, 252}, {57, 160},
					{110, 136}, {32, 37}, {149, 275}, {50, 76}, {50, 21}, {9, 245}, {133, 190}, {40, 153}, {284, 81}, {117, 191}, {76, 267}, {14, 207}, {216, 276}, {295, 1}, {65, 23}, {179, 108}, {44, 3},
					{100, 272}, {73, 275}, {198, 22}, {107, 293}, {198, 242}, {173, 70}, {81, 212}, {17, 143}, {237, 274}, {42, 260}, {15, 189}, {75, 42}, {5, 185}, {268, 289}, {10, 197}, {177, 55}, {162, 70},
					{179, 144}, {13, 25}, {259, 7}, {33, 111}, {79, 88}, {111, 87}, {255, 156}, {204, 21}, {94, 141}, {76, 289}, {227, 171}, {9, 87}, {102, 221}, {259, 267}, {228, 210}, {290, 81}, {63, 175},
					{52, 192}, {61, 26}, {291, 153}, {150, 275}, {211, 275}, {229, 5}, {42, 215}, {102, 278}, {113, 157}, {146, 212}, {207, 7}, {121, 288}, {207, 36}, {83, 261}, {64, 274}, {235, 146},
					{200, 58}, {45, 168}, {264, 273}, {204, 181}, {248, 42}, {182, 98}, {69, 126}, {219, 272}, {228, 119}, {98, 105}, {193, 49}, {19, 5}, {37, 34}, {130, 223}, {104, 196}, {137, 18},
					{94, 196}, {203, 120}, {273, 97}, {259, 86}, {77, 256}, {28, 256}, {284, 291}, {266, 164}, {251, 226}, {127, 159}, {98, 47}, {225, 197}, {279, 222}, {58, 125}, {44, 15}, {251, 51},
					{259, 58}, {55, 67}, {151, 93}, {181, 265}, {234, 47}, {155, 38}, {279, 88}, {183, 91}, {177, 203}, {89, 52}, {231, 191}, {9, 85}, {282, 242}, {128, 283}, {261, 171}, {248, 109},
					{247, 4}, {89, 271}, {101, 140}, {145, 213}, {62, 138}, {279, 269}, {92, 136}, {273, 45}, {132, 90}, {202, 149}, {85, 65}, {53, 113}, {277, 31}, {119, 150}, {253, 192}, {181, 136},
					{245, 297}, {214, 294}, {261, 266}, {186, 22}, {41, 225}, {233, 187}, {139, 54}, {194, 86}, {131, 220}, {45, 241}, {173, 280}, {289, 205}, {151, 124}, {107, 219}, {238, 53}, {238, 280},
					{88, 202}, {161, 201}, {289, 54}, {94, 129}, {35, 212}, {151, 38}, {144, 67}, {169, 236}, {264, 148}, {76, 241}, {35, 97}, {214, 278}, {99, 224}, {217, 6}, {57, 67}, {161, 134},
					{152, 241}, {118, 196}, {285, 22}, {93, 284}, {148, 34}, {231, 157}, {181, 105}, {81, 212}, {91, 95}, {55, 185}, {248, 55}, {61, 207}, {247, 62}, {184, 4}, {141, 256}, {194, 98},
					{198, 218}, {258, 175}, {145, 213}, {19, 254}, {104, 160}, {281, 168}, {205, 97}, {143, 46}, {234, 170}, {92, 206}, {131, 72}, {234, 210}, {183, 73}, {237, 84}, {139, 138},
					{285, 174}, {171, 160}, {55, 264}, {231, 128}, {75, 280}, {282, 278}, {119, 156}, {258, 131}, {21, 126}, {19, 142}, {293, 68}, {169, 72}, {70, 272}, {9, 216}, {243, 190},
					{95, 130}, {104, 3}, {119, 196}, {182, 207}, {293, 69}, {282, 50}, {182, 73}, {83, 179}, {132, 260}, {85, 272}, {182, 140}, {243, 34}, {227, 238}, {2, 146}, {55, 149}, {185, 82},
					{119, 159}, {65, 110}, {177, 13}, {151, 125}, {199, 206}, {161, 123}, {51, 17}, {13, 96}, {122, 79}, {57, 114}, {161, 29}, {234, 278}, {244, 226}, {86, 214}, {177, 233}, {223, 275},
					{183, 3}, {161, 129}, {61, 109}, {186, 97}, {261, 8}, {259, 123}, {261, 166}, {208, 6}, {91, 32}, {66, 228}, {263, 106}, {261, 118}, {163, 207}, {71, 56}, {287, 268}, {221, 225}, {115, 196},
					{12, 231}, {20, 6}, {263, 184}, {13, 231}, {95, 209}, {24, 196}, {121, 5}, {249, 164}, {244, 270}, {15, 41}, {297, 171}, {110, 179}, {41, 36}, {77, 3}, {214, 243}, {167, 252}, {141, 111},
					{238, 20}, {217, 102}, {218, 18}, {55, 282}, {69, 96}, {227, 109}, {132, 293}, {281, 240}, {127, 232}, {182, 130}, {88, 123}, {182, 114}, {188, 167}, {58, 20}, {236, 135}, {69, 70},
					{12, 72}, {80, 191}, {199, 257}, {57, 188}, {227, 211}, {247, 9}, {213, 276}, {44, 175}, {177, 10}, {256, 220}, {237, 289}, {155, 92}, {203, 197}, {20, 153}, {71, 5}, {285, 112},
					{75, 130}, {245, 294}, {146, 188}, {19, 59}, {19, 241}, {119, 252}, {8, 187}, {231, 233}, {259, 44}, {293, 212}, {264, 92}, {50, 163}, {123, 94}, {27, 14}, {258, 16}, {13, 279},
					{85, 153}, {169, 7}, {14, 25}, {150, 213}, {260, 16}, {162, 90}, {293, 158}, {267, 82}, {238, 100}, {161, 21}, {233, 174}, {296, 289}, {259, 296}, {56, 223}, {145, 187}, {18, 16},
					{186, 128}, {193, 298}, {296, 161}, {77, 162}, {290, 7}, {128, 34}, {55, 152}, {50, 236}, {284, 155}, {117, 49}, {151, 93}, {85, 11}, {124, 126}, {271, 90}, {93, 246}, {239, 213},
					{200, 216}, {166, 267}, {243, 144}, {263, 270}, {15, 103}, {177, 278}, {182, 85}, {168, 250}, {189, 163}, {189, 233}, {258, 256}, {58, 192}, {251, 171}, {81, 101}, {43, 151}, {106, 290},
					{161, 145}, {147, 3}, {251, 87}, {194, 149}, {31, 116}, {211, 263}, {121, 87}, {144, 262}, {268, 177}, {204, 211}, {82, 240}, {233, 245}, {246, 202}, {218, 166}, {71, 78}, {293, 285},
					{145, 258}, {103, 239}, {110, 200}, {282, 162}, {21, 40}, {183, 253}, {50, 292}, {66, 157}, {60, 22}, {33, 187}, {227, 263}, {138, 170}, {70, 72}, {161, 281}, {137, 26}, {11, 127},
					{246, 58}, {88, 53}, {33, 8}, {86, 140}, {188, 26}, {129, 214}, {21, 109}, {47, 280}, {110, 117}, {52, 192}, {94, 153}, {139, 70}, {229, 97}, {19, 77}, {106, 140}, {137, 84}, {189, 48},
					{148, 218}, {124, 278}, {214, 224}, {149, 56}, {239, 20}, {294, 4}, {290, 284}, {169, 185}, {124, 120}, {15, 185}, {63, 232}, {77, 32}, {217, 102}, {183, 109}, {247, 113}, {188, 37},
					{237, 197}, {2, 131}, {115, 210}},
			},
			want: []int{61, 64, 234},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toposortKahnAlgo(tt.args.k, tt.args.con); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toposortKahnAlgo() = %v, want %v", got, tt.want)
			}
		})
	}
}
