package main

import (
	"fmt"
	"hash/crc32"
	"os"
	"io"
)

func getHash(fileName string) (uint32, error){
	// open the file

	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}

	defer f.Close()
	h := crc32.NewIEEE()
	// copy file to into the hasher
	bytesNo, err := io.Copy(h, f)
	fmt.Println("Number of bytes written" , bytesNo)

	// We don't care about how many bytes were written, but we do want to handle error

	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("test1.txt")
	if err != nil {
		return
	}

	h2, err := getHash("test2.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}

