package main

import (
	"fmt"
	"math/rand"
	"time"
)

const ALPHABET = "0123456789ABCDEFGHIJKLMNOPQRSTUVZ"

func mul(a uint32, b uint32) (result uint32, overflow uint32) {
	total := uint64(a) * uint64(b)
	return uint32(total), uint32(total >> 32)
}

/* Helper function to calculate magic from loop counter */
func calcMagic(counter int32) uint32 {
	_, o := mul(0x38e38e39, uint32(counter))
	return uint32(counter) - ((((o >> 3) << 3) + (o >> 3)) << 2)
}

/* Reversed Ord function, used to generate keys */
func reverseOrd(ord uint32) byte {
	switch {
	case ord < 0xa:
		return byte(ord + 0x30)
	case ord >= 0xa && ord <= 0x23:
		return byte(ord + 0x37)
	}
	panic(fmt.Sprintf("Invalid ord %x\n", ord))
}

func keygen() string {
	key := "0FUN" // Brand the key :)
	// Generate the first 15-len("COCOZZA") chars of the key
	for i := 0; i < 15-len("0FUN"); i++ {
		key += string(ALPHABET[rand.Int()%len(ALPHABET)])
	}
	// Calculate the counter for the random 15 bit key
	counter := int32(0)
	for i := int32(0); i < int32(len(key)); i++ {
		counter += (ord(key[i]) + 1) * (i + 1)
	}
	// Calculate the magic
	magic := calcMagic(counter)
	// Get the final character and append it to the key
	key += string(reverseOrd(magic))
	return key
}

func ord(char byte) int32 {
	switch {
	case '0' <= char && char <= '9':
		return int32(char) - 0x30
	case 'A' <= char && char <= 'Z':
		return int32(char) - 0x37
	}
	panic("Found Invalid Character!")
}

func validateKey(key string) bool {
	if len(key) != 16 {
		return false
	}

	counter := int32(0)
	for i := int32(0); i < int32(len(key)-1); i++ {
		counter += (ord(key[i]) + 1) * (i + 1)
	}
	magic := calcMagic(counter)

	return uint32(ord(key[len(key)-1])) == magic
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		key := keygen()
		fmt.Println(key)
	}
}
