// Example1: 
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {

	//s1 := []byte{0x48, 0x28, 0x2F, 0x49, 0x30, 0x2E, 0x48, 0x28, 0x2F, 0x49, 0x30, 0x2E, 0x48, 0x28, 0x2F, 0x49}
	//s1 = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16}

	buf := new(bytes.Buffer)
	var i uint16 = 0x48
	pp.Println("i=", i)

	err := binary.Write(buf, binary.BigEndian, i)

	if err != nil {
		fmt.Println("error ....=", err.Error())
	}

	reader := bytes.NewReader(buf.Bytes())

	type rData struct {
		Val [2]byte
	}

	rd := &rData{}
	var j uint16

	if err := binary.Read(reader, binary.LittleEndian, &j); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(rd.Val[0] & 0xF)

	fmt.Println(rd.Val)

	fmt.Printf("%b", rd.Val)

	fmt.Println(j)

}

// output
/*
"i=" 0x0048
0
[0 0]
[0 0]18432
*/


// Example2: 
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {

	//s1 := []byte{0x48, 0x28, 0x2F, 0x49, 0x30, 0x2E, 0x48, 0x28, 0x2F, 0x49, 0x30, 0x2E, 0x48, 0x28, 0x2F, 0x49}
	//s1 = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16}

	buf := new(bytes.Buffer)
	var i uint16 = 0x48
	pp.Println("i=", i)

	err := binary.Write(buf, binary.LittleEndian, i)

	if err != nil {
		fmt.Println("error ....=", err.Error())
	}

	reader := bytes.NewReader(buf.Bytes())

	type rData struct {
		Val [2]byte
	}

	rd := &rData{}
	if err := binary.Read(reader, binary.BigEndian, rd); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(rd.Val[0] & 0xF)

	pp.Println(rd.Val)

	fmt.Printf("%b", rd.Val)
}


// output
/*
"i=" 0x0048
8
[2]uint8{
  0x48, 0x00,
}
[1001000 0]
*/


// Example3:
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {

	//s1 := []byte{0x48, 0x28, 0x2F, 0x49, 0x30, 0x2E, 0x48, 0x28, 0x2F, 0x49, 0x30, 0x2E, 0x48, 0x28, 0x2F, 0x49}
	//s1 = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16}

	buf := new(bytes.Buffer)
	var i uint32 = 16909060
	var j uint32 = 67305985

	err := binary.Write(buf, binary.BigEndian, i)
	err = binary.Write(buf, binary.BigEndian, j)
	fmt.Println("Buffer for W")
	pp.Println("i", i)
	pp.Println("j", j)
	if err != nil {
		fmt.Println("error ....=", err.Error())
	}

	reader := bytes.NewReader(buf.Bytes())

	type rData struct {
		Val [8]byte
	}

	rd := &rData{}
	if err := binary.Read(reader, binary.LittleEndian, rd); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	fmt.Println(rd.Val[0] & 0xF)
	fmt.Println(rd.Val[1] >> 4)
	fmt.Println(rd.Val[2] & 0xF)
	fmt.Println(rd.Val[2] | 0xF)
	fmt.Println(rd.Val[5] | 0xF)

	pp.Println(rd.Val)
}
//output
/*
Buffer for W
"i" 0x01020304
"j" 0x04030201
1
0
3
15
15
[8]uint8{
  0x01, 0x02, 0x03, 0x04, 0x04, 0x03, 0x02, 0x01,
}

*/
