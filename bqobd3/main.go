package main

import "fmt"

/* I call this Extended Fibonacci :'D */
func calc(input uint32) uint32 {
	v0 := uint32(0x2345)
	v1 := uint32(1 + 0x2345)
	v2 := uint32(4 + 0x2345)
	v3 := uint32(9 + 0x2345)
	v4 := uint32(16 + 0x2345)

	for i := uint32(4); i < input; i++ {
		v0, v1, v2, v3, v4 = v1, v2, v3, v4, (v0*uint32(0x1234))+(v4-v3)+(v2-v1)
	}

	return v4
}

func main() {
	fmt.Printf("0x%x\n", calc(0x18e28))
}
