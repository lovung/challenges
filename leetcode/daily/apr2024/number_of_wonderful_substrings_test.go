package apr2024

import (
	"testing"
)

func Test_wonderfulSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				word: "cdaabbadac",
			},
			want: 29,
		},
		{
			args: args{
				word: "aabb",
			},
			want: 9,
		},
		{
			args: args{
				word: "bbgfeijhhjghbbageiciicfegdjedjebigbiajhibcadigjifbaghgfbceggbfcaehhjechhfjghbchgafjebaefcfibddgghcaibhbdejdhbfcbbecdfgiffcbhgaeaceijchfeedigcbibibhiechfhajgddccefjcfhjgaajddgcigdegjibhgihbbadaheggjiedddgfafdijcfffcbdaiejiheghgffdbgchiciiaieijcehbdceagejbidffhdchahhajhhfagcfdjchhhhdjfjhghbhjafdhhcjgbjbcaidfeeebgefbccaebefdejcjbdjaibhceiegfffhhaeaejfjbiheigbjjjhaajhcffhhbgdifgajhcfabiegidggfdfhcffcabiddaaegihegdbcaebhfjcbdefdbddeeagjaehbcfjbhfafebfhcefiijdacjjcijahajbejbhagaeaiieidbibhfagdgicgjeeddeegieiecfbghcfbijfdifdhiadhdcbadceeciieffdjbeehebbgaefjjfbhidjhigaebdhfbfjhdajhdfdeggaccdihaegchjdcfbiggeecjfjhjeiifaafcccegfffigbiegfijeijebaffghfiaaahjijbjigcijehahbcaiebhadbaibceedaacggahbgehghgabhhdbiaagebcjabgfhjjidhabhfehefdchdffhfjidjefhfeiijbfgjfggjgiijdjjiifhehcjchgjcejghchhibbdgfhhaaghcajighedbjgefbgabfdecfdbaadbbbdiddibjaddhefgiaifgbjbcdgafbbiddiffhebgeaaeicaefdagjjfcaeeibiicacdgiacbhaiahicjbcefgjhajfjfagbifjhdjcbafccadhecfdghhaacjjjijfeeicgabiiifhcaijceiibfjdebifghjfgcibdbffidbdicgiahbccecdhaefjhhhhaiejhggiigjfahbdddfcgbbdgfjgabhhggcaagcgbajabcaebcjfiffhjeadhhjfgjgggjbjideeegffcbaegicifhbbageddfhiccbjaedhibcgbefgdcaffihijfgabecfchcbddjaaeggaeahibjaeaidcbeccjadbhfiaagadjcabdefabdefgdfgaihgjeejdaiagideihghdchjhfieajcajghcjjceiiedidegjgeihdeaaebcaieaaibdhefababigeddfjbjacbgbjdcgcacgdbbecbjgebjjbcihecjejbececfidafadehaffhfdhbcijjhjhhajchdhdgbjgifadfaffciacehdfcifjhafcghegijfbijidgefegfcbabjhiifjibfgbijbifdeaghgbadfcfjbbgiedbfigahfgdcejbgeeafbgehcciihiighchjjafcedhdjibhhgfbdhcbfbhgjidiheejechaibdeeghfiifagdfbgcbgfhcbaafdjjdfacijfghagjdcfaedgjfdjaiadchbhjfijiidadccbihbbdiebffjdaddjicfgadjhegeegjdeghhhfejigaghbjafhjgdgejdeedfcacehdigacbdfieahifacibiefadjjddjachhaefdgfggcecfdjbejbchachchcgjaggbiddchgajhjbieihbjdbbadhcfjabdjecdfdggjhffhfefeffgjaifdjbgbjfhcecdcicdhfcfgifhhiddhfifgafdbbecehdeedefjedgcgdbbgiigdccbcidjfjbjiijdbagebidbeifjeadbifaahfcbcgjfbbidgfejghbjffaaajhcbcgbejdhdaifdgafbcbbabgbicdbagdeecahfcgdbehdbjbagbfahageajeddhhbfhdahfjbjcjdcafgjjaacidbjddcabhjgeghbchecdiieccciighicafjajijeafchfjaghjhjfajdaicejbhgfbdgfbdcfbjdgeeiecgdbeabdjijdbbecffbbibceiefhhjahjcfchafdffhbiijjafffeiejjgbdgficecaibgagjjddeefdaacacgjgechacdiagidbcidefahfijbbjciaabifgjcgfbaahjhajhbhiefdhhbhcgbcgbdfdciaddedbfegbdeedgbhgicbfeiaiiiidffihifffdaibabggfjhfadddhcbegeefffcfefbjcfafhjajhjbcjhhigjccfaabeecdedhejjdgeaedgfiabheccadegeidejechgjahbdicfiidecidebjehebaibdafacfaadbfcfbbfjehahfgbhjfhabbigfgfaicfjbgfabcghedchgabiiiibdieegeccacagegbcegijjfeffdghhcfcaegdgdjidihhgggjdacjdjebhihhgbchfhfcifgcgabdcihcfeebceejjabaaajafiagadbgajfbbiffihjagihhjcbjjjgdfjeifegaahhjhgiachhedgjfhihcdedbbfffaigiihafcjejfhghefafjbaeebbeegbbjgecaeefcidhfhghjjcdjeiajbcbcahaedahjeifjgicjagicdiaeeghbhdighbhhijjheggaidgggehcidedhjbadidafjfacfjgcecgbbfhaiabhdfdddfeiidhjfigeejjiagcijbgjbcdabibfdgjdghdcbigcbagebdaiaddccchhibdheeffdfifjjigidjbbcidajhfgibibcceaibfebfgfjjabdeihbijdjacdaaijhacddggcfeggccgbfahjecaicbiabiaieaicjeehbfddedfjhieeacffgbehdfbhhbiaigghifabadbjbgigjhafhjaicjajgdbcejcggfcedgedeideaacijgicdfajjhfcejbadgebdabaacchghifefcajbgchbcjgecgcahgbihddgeigjidegjiedfddgjdfbhdbgeceiaibfihegfifgdjehdcejfefdbbgciicidagefhidfahehidghchbjccjdbbgiaggajecgidbbfjbfaegdhejicchffhdgbddifbaicfafdfcchgacihhbigfegaabbacfhgbbfghifjdfgebggehcafghhjffdcfjiefgafijaibhhdiehcichiggaefjbbejcgdfheaaaaifgcbfhjhghbebbbacafjcehcidbiccejiecgjejigicbajedihhadaidabiheacffaafjeceiadefbbgghfgehebdcecfiifjajcgfgafbjbehhbfhcjdachdegfggihbgcdfegfhiecaiijjhjdjaehfejjgjhahbfagchdefhghgfccihgacafejcaadeffcfdagjfcbicgbadgdjjghdeihgcfecdadigihhffdjahffachjgddcgcajhedgedgjdfaiaedhhgidggdfifhbhbccfjijjijbaebcjebfdeeeihhgaccdjegejfbgahcaaahehgfgahbaifddifafdhddjjigehagfjaigiiffbiidijjbcbigcgfifihicgeddjegbbjbgfeeegdghhaceadbedcadfccfdcfbegeiiedibhcheigijciihihiehaccdaeggbibafhfcjjgdiffcjfeeiihdbfcedhbgcajhjddgfifjccheaibeaacijfibdhiefifdfeigefjfcjfdbbcgbjicigigebggdhgieabigbacfafehifbjgcefdjdefjcjgbifieeeghbijjiibiaajhejaaajfbiihfhidiedcfeahjhdghjjgjebagdcdfidihcgccihaejafajbfejjjh",
			},
			want: 89065,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderfulSubstrings(tt.args.word); got != tt.want {
				t.Errorf("wonderfulSubstrings() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderfulSubstrings1(tt.args.word); got != tt.want {
				t.Errorf("wonderfulSubstrings1() = %v, want %v", got, tt.want)
			}
		})
	}
}
