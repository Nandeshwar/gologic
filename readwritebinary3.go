package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {

	s1 := []byte{0x48, 0x28}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, s1)
	if err != nil {
		fmt.Println("error ....=", err.Error())
	}

	reader := bytes.NewReader(buf.Bytes())

	type rData struct {
		Val [2]byte
	}

	rd := &rData{}
	if err := binary.Read(reader, binary.LittleEndian, rd); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(rd.Val[0] & 0xF)
	fmt.Println(rd.Val[1] >> 5)
}
