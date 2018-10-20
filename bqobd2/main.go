package main

import "fmt"

/* Actually implemented function */
func rfib(input uint32) (output uint32) {
	val := input
	if val <= 1 {
		output = val
	} else {
		output = rfib(val-1)+rfib(val-2)
	}
	return
}

/* "Optimized" function */
func fib(n uint32) uint32 {
  a := uint32(1)
  b := uint32(1)
  for i := uint32(2); i < n; i++ {
    a, b = b, a+b
  }
  return b
}

func main() {
  fmt.Printf("0x%x\n", fib(1015));
}
