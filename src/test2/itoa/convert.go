package convert

import (
   "fmt"
   "strconv"
)

func convert1(i int) string {
   s := fmt.Sprintf("%d", i)
   return s
}

func convert2(i int) string {
   s := strconv.Itoa(i)
   return s
}

func convert3(i int) string {
   s := strconv.FormatInt(int64(i), 10)
   return s
}
