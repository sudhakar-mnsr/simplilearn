package convert

import "testing"

var result string

func BenchmarkConvert1(b *testing.B) {
   var s string
   for n := 0; n < b.N; n++ {
      s = convert1(n)
   }
   result = s
}

func BenchmarkConvert2(b *testing.B) {
   var s string
   for n := 0; n < b.N; n++ {
      s = convert2(n)
   }
   result = s
}

func BenchmarkConvert3(b *testing.B) {
   var s string
   for n := 0; n < b.N; n++ {
      s = convert3(n)
   }
   result = s
}
