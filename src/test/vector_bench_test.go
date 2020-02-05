package vector

import (
   "math/rand"
   "testing"
   "time"
)

func BenchmarkVectorEqual1(b *testing.B) {
   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   for i := 0; i < b.N; i++ {
      v1 := New(r.Float64(), r.Float64())
      v1.Eq(v1)
   }
}

func BenchmarkVectorEqual2(b *testing.B) {
   r := rand.New(rand.NewSource(time.Now().UnixNano()))
   for i := 0; i < b.N; i++ {
      v1 := New(r.Float64(), r.Float64())
      v1.Eq2(v1)
   }
}
