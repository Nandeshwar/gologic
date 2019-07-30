/*
Encode logic:
input: 0x12, 0x34, 0x5F, 0x9A, 0x0F
output: 74565, -10.2, N

How was this hex encoded
----------------------
weather pst Id : 74565 ---Number , Max can be 100, 000               -> 20 bits,  100, 000 = 0001 1000 0110 1010 0000
Temp: (C):  (-99.9 -> 99. 99) : -99.9 * 10, 99.9 *10 => -999 -> 999 : 12 bits,  999      = 0111 1110 0111
Wind Direction: "N", "E", "W", "S", "NE", "NW", "SE", "SW"
                 0    1    2    3    4     5     6     7
Wind dir: number (0 - 15) :                                           4 bits : 15        = 1111

Weather_Post {
  postId : 74565  -- 20 bits -> 0001 0010 0011 0100 0101
                              Represent in 20 bits:  0001 0010 0011 0100 0101
  temp: -10.2  -- 12 bits ->
         - : To get - 10.2
   74565:  0001 0010 0011 0100 0101
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
0001 0010 0011 0100 0101 1111 1001 1010 0000 1111
|                      | |            | |  |
+----------------------+ +------------+ +--+
0001 0010 = 0x12
0011 0100 = 0x34
0101 1111 = 0x5F
1001 1010 = 0x9A
0000 1111 = 0x0F

Encoded hex: 12 34 5F 9A 0F
Total bits required: 36 bits
Multiple of 8 : 40 bits : 4 byte can hold all this info

Program given below is to decode windDir: 102
How to decode?

1. Store 1st 3 bytes in 3 int32 variable
a0: 0001 0010 in 4 bytes int32
a1: 0011 0100 in 4 bytes int32
a3: 0101 1111 in 4 bytes int32

Expected binary for 74565 in 4 bytes: 0000 0000 0000 0001 0010 0011 0100 0101

How to get?
Method1:
a0 in 4 bytes: 00000000  00000000 00000000 00010010
a0: left shift << by 12 bits:
    -> 00000000 00000001 00100000 00000000 : Compare to expected binary above. 0001 and 0010 are in right position

a1 in 4 bytes: 0000 0000 0000 0000 0000 0000 0011 0100
a1: right shift << by 4 bits:
0000 0000 0000 0000 0000 0011 0100 0000 : Compare to expected binary above. 0011 and 0100 are in right position

a3 in 4 bytes: 00000000  00000000 00000000 01011111
a3: right shift >> by 4 bits:
  00000000  00000000 00000000 00000101: Compare to expected binary above. 0101 is in right position

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

	s1 := [5]byte{0x12, 0x34, 0x5F, 0x9A, 0x0F}

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
		// weather post id
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
output:
[5]uint8{
  0x12, 0x34, 0x5f, 0x9a, 0x0f,
}
74565
*/
