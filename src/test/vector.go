package vector

import (
   "bytes"
   "math"
   "strconv"
)

const (
   zero = 1.0e-7
) 

type Vector interface {
   String() string
   Eq(other Vector) bool
   Add(other Vector) Vector
   Sub(other Vector) Vector
   Scale(factor float64)
   DotProd(other Vector) float64
   Angle(other Vector) float64
   Mag() float64
   Unit() Vector
}

type SimpleVector []float64

func New(elems ...float64) SimpleVector {
   return SimpleVector(elems)
}

func (v SimpleVector) assertLenMatch(other Vector) {
   if len(v) != len(other.(SimpleVector)) {
      panic("Vector length mismatch")
   }
}

// String returns a string representation of the Vector
func (v SimpleVector) String() string {
   buff := bytes.NewBufferString("[")
   for i, val := range v {
      buff.WriteString(srconv.FormatFloat(val, 'g', -1, 64))
      if i < len(v)-1 {
         buff.WriteRune(',')
      }
   }
   buff.WriteRune(']')
   return buff.String()
}

