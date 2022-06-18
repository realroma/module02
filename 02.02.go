
package main

import (
   "fmt"
  "math"
   ) 
var (
   pi float64 = 3.14
   a int = 35
   r float64
   R *float64 = &r 
)

func main() {
   r = float64(a)/2/pi
   fmt.Printf( "%.2f радиус", *R)
   S(r) 
}
func S (r float64) {
   r = math.Pow(r, 2)
   x := 2 * pi * *R
   fmt.Printf("\n%.2f площадь", x)
}
