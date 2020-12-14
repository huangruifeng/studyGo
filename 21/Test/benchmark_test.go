package functions

import (
	"bytes"
	"testing"
)

func getParam() []string {
	return []string{"100asdadasdasdq",
		"200qweqeqeasdad",
		"300adsadasdadad",
		"400asdadadadad",
		"123123asdasasdad",
		"asdadasdasdads",
		"123132qsdasdasdads",
		"123asdasdasd "}
}

// go test -bench=. -benchmem
//win go test -bench="."" -benchmem

func BenchmarkStringaddTest(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var result string
		for _, s := range getParam() {
			result = result + s
		}
	}
	b.StopTimer()
}

func BenchmarkByteBufferTest(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var res bytes.Buffer
		for _, s := range getParam() {
			res.WriteString(s)
		}
	}
	b.StopTimer()
}
