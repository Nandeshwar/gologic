package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// Create a hasher
	h := crc32.NewIEEE()
	//write out data to it
	h.Write([]byte("test"))
	// calculate the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)

	// Create a hasher
	h2 := crc32.NewIEEE()
	//write out data to it
	h2.Write([]byte("test"))
	// calculate the crc32 checksum
	v2 := h.Sum32()
	fmt.Println(v2)


}