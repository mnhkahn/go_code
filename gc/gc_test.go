package gc

import "testing"

var payafterids = []int64{10112959721, 10048537219, 10113119386, 10119724841, 10116503893, 10156167482, 10114519415, 10112644839, 10122313331, 10121334423, 10113769024, 10112946259, 10113935097, 10156558374, 10119733312, 10115877513, 2943852, 10105981375, 10113743549, 10087919754, 10115848417, 10116008447, 10120047563, 10116520654, 10115483651, 10207268890, 10115172705, 10112158424, 10113891192, 10112953497, 10110105882, 10126036196, 10149239840, 10124528038, 10117202713, 10119178201, 10110070945, 10115873894, 10113903511, 10113918112, 10118425581, 10115029660, 10162798393, 10113300456, 10377621715, 2930900, 10115966725, 10165823684}
var ids = []int64{10116503893, 10105981375, 10124528038, 10115029660, 10120101518, 2943766, 10118425581, 10209227477, 10126490232, 10158321072, 10213313050, 10115261583, 10147007432, 10115338243, 2943386, 2919748, 10113290635, 2943376, 2812129, 3029250, 10242692188, 2943880, 3014478, 3029416, 10380295975, 2811967, 10195762839, 10393710036, 10397100332, 2815871, 2943826, 10380797788, 10394372679, 2856261, 2937894, 2931148, 3014332, 10374479747, 2810119, 10358314023, 10365953479, 2931158, 2943394, 10395101981, 2800075, 10298663619, 2930900, 2937212, 2935526, 10119722690, 10394139779, 2813991, 3028944, 10110917581, 10324180068, 10127255569, 10121291949, 10113698738, 10185863750, 10126479221}

func TestPayAfter(t *testing.T) {
	if PayAfter(ids, payafterids) != PayAfter2(ids, payafterids) {
		t.Error("PayAfter Error")
		t.Log(PayAfter(ids, payafterids))
		t.Log(PayAfter2(ids, payafterids))
	}
	if PayAfter(ids, payafterids) != PayAfter22(ids, payafterids) {
		t.Error("PayAfter Error3")
		t.Log(PayAfter(ids, payafterids))
		t.Log(PayAfter22(ids, payafterids))
	}
}

func BenchmarkPayAfter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PayAfter(ids, payafterids)
	}
}

func BenchmarkPayAfter2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PayAfter22(ids, payafterids)
	}
}

// func BenchmarkPayAfter3(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		PayAfter3(ids, payafterids)
// 	}
// }

// func BenchmarkFmtSprintf(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		z := FmtSprintf(i)
// 		_ = z
// 	}
// }

// func BenchmarkItoa(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		z := Strconv(i)
// 		_ = z
// 	}
// }

// func BenchmarkFormat(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		z := FormatInt(i)
// 		_ = z
// 	}
// }

// func BenchmarkNewIds(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		z := NewIds()
// 		_ = z
// 	}
// }

// func BenchmarkNewIdsPool(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		z := NewIdsPool()
// 		_ = z
// 	}
// }

// func BenchmarkPrintln(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Println(i)
// 	}
// }

// func BenchmarkFmtPrintln(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		FmtPrintln(i)
// 	}
// }

// func BenchmarkLogPrintln(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		LogPrintln(i)
// 	}
// }

// func BenchmarkBeegoInfo(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		BeegoLog(i)
// 	}
// }

// func BenchmarkAddAttrToMap2(b *testing.B) {
// 	res := map[string]int64{"870_1111": 1}
// 	exclude := map[int64]int8{111: 1}
// 	for i := 0; i < b.N; i++ {
// 		temp := map[string]int64{"870_0": 1}
// 		res = AddAttrToMap2(res, temp, exclude)
// 	}
// }
