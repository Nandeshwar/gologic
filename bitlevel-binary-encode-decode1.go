/*
Encode logic:
input: 00 06 6F 9A 0f
output: 102, -10.2, N

How was this hex encoded
----------------------
weather pst Id : 102 ---Number , Max can be 100, 000               -> 20 bits,  100, 000 = 0001 1000 0110 1010 0000
Temp: (C):  (-99.9 -> 99. 99) : -99.9 * 10, 99.9 *10 => -999 -> 999 : 12 bits,  999      = 0111 1110 0111
Wind Direction: "N", "E", "W", "S", "NE", "NW", "SE", "SW"
                 0    1    2    3    4     5     6     7
Wind dir: number (0 - 15) :                                           4 bits : 15        = 1111

Weather_Post {
  postId : 102  -- 20 bits -> 0110 0110
                              Represent in 20 bits: 0000 0000 0000 0110 0110
  temp: -10.2  -- 12 bits ->
         - : To get - 10.2
   102:  0110 0110
   Not:  1001 1001
  Add 1: 1001 1010

  Represent it in 12 bits :
    1111 1001 1010
    ^
    |---------: first one is unused

  windDir: N    -- 4 bits : N is 0 as per map given above:
                0000
    -- Represent in 1 byte: 0000 1111
         Last 4 bits unused.

Questions:
   Why first 4 bit is unused in case of weather?
   Why last 4 bit is unused in case of windDir

}

A full 40 bits: 8 byte decoded data
0000 0000 0000 0110 0110 1111 1001 1010 0000 1111
|                     |   |          |   |    |
+---------------------+   +----------+   +----+
0000 0000 = 0x00
0000 0110 = 0x06
0110 1111 = 0x6F
1001 1010 = 0x9A
0000 1111 = 0x0F

Encoded hex: 00 06 6F 9A 0F
Total bits required: 36 bits
Multiple of 8 : 40 bits : 4 byte can hold all this info

Program given below is to decode windDir: 102
How to decode?

1. Store 1st 3 bytes in 3 int32 variable
a0: 0000 0000 in 4 bytes int32
a1: 0000 0110 in 4 bytes int32
a3: 0110 1111 in 4 bytes int32

Expected binary for 102 in 4 bytes: 0000 0000 0000 0000 0000 0000 0110 0110

How to get?
Method1:
a0 : left shift << by 12 bits
a1 : left shift << by 4 bits:
a3: right shift >> by 4 bits:

Method2:

*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/k0kubun/pp"
)

func main() {

	s1 := [5]byte{0x00, 0x06, 0x6F, 0x9A, 0x0F}

	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, s1)
	if err != nil {
		fmt.Println("error ....=", err.Error())
	}

	reader := bytes.NewReader(buf.Bytes())

	type rData struct {
		Val [5]byte
	}

	rd := &rData{}
	if err := binary.Read(reader, binary.LittleEndian, rd); err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	pp.Println(rd.Val)

	var (
		//post id
		a0 int32
		a1 int32
		a2 int32
	)

	a0 = int32(rd.Val[0]) << 12
	a1 = int32(rd.Val[1]) << 4
	a2 = int32(rd.Val[2]) >> 4

	id := a0 + a1 + a2
	fmt.Println(id)

}

/*
 Output:
 [5]uint8{
  0x00, 0x06, 0x6f, 0x9a, 0x0f,
}
102
*/
