package main

// NOTE to self: this code fucking sucks.

import (
	"fmt"
	"math/rand"
	"time"
)

const ALPHABET = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

/* Reversed Ord function, used to generate keys */
func reverseOrd(ord int32) byte {
	switch {
	case ord >= 0 && ord <= 9:
		return byte(ord + 0x30)
	case ord >= 0xa && ord <= 0x23:
		return byte(ord + 0x37)
	default:
		return reverseOrd((0x24 + ord) % 0x24)
	}
}

func keygen() string {
	key := []byte("0000000000000000") // init an empty key
	// Get some char randomly
	key[0] = ALPHABET[rand.Int()%(0x24-6)] // -6 to prevent bugs on constraint 3
	key[8] = ALPHABET[rand.Int()%0x24]
	key[0xb] = ALPHABET[rand.Int()%0x24]
	key[0xd] = ALPHABET[rand.Int()%0x24]
	// Calculate the rest
	key[1] = reverseOrd(0xe - ord(key[0]))
	key[2] = reverseOrd(0x6 + ord(key[0]))
	key[3] = reverseOrd(0x18 - ord(key[2]))
	key[5] = reverseOrd(0x4 - ord(key[3]) - ord(key[1]))
	key[4] = reverseOrd(0x16 - ord(key[5]) - ord(key[3]))
	key[6] = reverseOrd(0xd - ord(key[4]) - ord(key[2]))
	key[7] = reverseOrd(0x7 - ord(key[4]) - ord(key[1]))
	key[0xa] = reverseOrd(0x1f - ord(key[8]) - ord(key[6]))
	key[9] = reverseOrd(0x1b - ord(key[0xa]) - ord(key[8]))
	key[0xc] = reverseOrd(0x17 - ord(key[0xd]) - ord(key[0x7]))
	key[0xf] = reverseOrd(0x14 - ord(key[0xc]) - ord(key[9]))
	key[0xe] = reverseOrd(0xc - ord(key[0xf]) - ord(key[0xd]))
	return string(key)
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

func validateKey(key string) (valid bool, constraint int) {
	if len(key) != 16 {
		return false, -1
	}

	switch {
	// key_constraint_01
	case (ord(key[0])+ord(key[1]))%0x24 != 0xe:
		return false, 1
	// key_constraint_02
	case (ord(key[2])+ord(key[3]))%0x24 != 0x18:
		return false, 2
	// key_constraint_03
	case (ord(key[2])-ord(key[0]))%0x24 != 0x6:
		return false, 3
	// key_constraint_04
	case (ord(key[1])+ord(key[3])+ord(key[5]))%0x24 != 0x4:
		return false, 4
	// key_constraint_05
	case (ord(key[2])+ord(key[4])+ord(key[6]))%0x24 != 0xd:
		return false, 5
	// key_constraint_06
	case (ord(key[3])+ord(key[4])+ord(key[5]))%0x24 != 0x16:
		return false, 6
	// key_constraint_07
	case (ord(key[6])+ord(key[8])+ord(key[0xa]))%0x24 != 0x1f:
		return false, 7
	// key_constraint_08
	case (ord(key[1])+ord(key[4])+ord(key[7]))%0x24 != 0x7:
		return false, 8
	// key_constraint_09
	case (ord(key[9])+ord(key[0xc])+ord(key[0xf]))%0x24 != 0x14:
		return false, 9
	// key_constraint_10
	case (ord(key[0xd])+ord(key[0xe])+ord(key[0xf]))%0x24 != 0xc:
		return false, 10
	// key_constraint_11
	case (ord(key[8])+ord(key[9])+ord(key[0xa]))%0x24 != 0x1b:
		return false, 11
	// key_constraint_12
	case (ord(key[7])+ord(key[0xc])+ord(key[0xd]))%0x24 != 0x17:
		return false, 12
	}
	return true, 0
}

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		k := keygen()
		fmt.Println(k)
	}
}
