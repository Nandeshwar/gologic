package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	type data struct {
		Val1 uint8
		Val2 uint8
	}

	d := &data{}
	s1 := []byte{0x48, 0x28}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, s1)
	if err != nil {
		fmt.Println("error ....=", err.Error())
	}

	reader := bytes.NewReader(buf.Bytes())

	if err := binary.Read(reader, binary.LittleEndian, d); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(d.Val1 & 0xF)
	fmt.Println(d.Val2 >> 5)
}
