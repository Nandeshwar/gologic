package main

import (
	"fmt"
)

type Matrix struct {
	data       [][]int
	row        int
	column     int
	zeroRow    int
	zeroColumn int
}

func main() {
	m := Matrix{}
	m.getMatrix()
	m.Print()
	fmt.Println("Replace zero row column")
	m.findZeroRowColumn()
	m.setZeroRowColumn()
	m.Print()
}

func (m *Matrix) getMatrix() {
	var rowLen int
	var columnLen int
	fmt.Println("Enter Matrix row Length: ")
	fmt.Scanf("%d", &rowLen)

	fmt.Println("Enter Matrix colun Length: ")
	fmt.Scanf("%d", &columnLen)

	m.data = make([][]int, rowLen)
	m.row = rowLen
	m.column = columnLen

	fmt.Printf("Enter Matrix Elements")
	for i := 0; i < rowLen; i++ {
		m.data[i] = make([]int, columnLen)
		for j := 0; j < columnLen; j++ {
			fmt.Printf("\nEnter element: (%d,%d)\t", i, j)
			fmt.Scanf("%d", &m.data[i][j])
		}
	}

}

func (m *Matrix) Print() {
	fmt.Println("----Printing Matrix-----")
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.column; j++ {
			fmt.Printf("%d\t", m.data[i][j])
		}
		fmt.Println()
	}
}

func (m *Matrix) findZeroRowColumn() {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.column; j++ {
			if m.data[i][j] == 0 {
				m.zeroRow = i
				m.zeroColumn = j
			}
		}
	}
}

func (m Matrix) setZeroRowColumn() {
	for i := 0; i < m.row; i++ {
		for j := 0; j < m.column; j++ {
			if i == m.zeroRow || j == m.zeroColumn {
				m.data[i][j] = 0
			}
		}
	}
}

/*
output:
GOROOT=/usr/local/Cellar/go/1.12.5/libexec #gosetup
GOPATH=/Users/nandeshwar.sah/go #gosetup
/usr/local/Cellar/go/1.12.5/libexec/bin/go build -o /private/var/folders/nc/7bk5nj1x1db3_lg5_4fbyb4hnhxf8p/T/___go_build_MatrixEntrieRowColumnZero_go /Users/nandeshwar.sah/go/src/Prem/cmd/prem2/MatrixEntrieRowColumnZero.go #gosetup
/private/var/folders/nc/7bk5nj1x1db3_lg5_4fbyb4hnhxf8p/T/___go_build_MatrixEntrieRowColumnZero_go #gosetup
Enter Matrix row Length: 
3
Enter Matrix colun Length: 
3
Enter Matrix Elements
Enter element: (0,0)	1

Enter element: (0,1)	2

Enter element: (0,2)	3

Enter element: (1,0)	4

Enter element: (1,1)	5

Enter element: (1,2)	6

Enter element: (2,0)	7

Enter element: (2,1)	0

Enter element: (2,2)	8
----Printing Matrix-----
1	2	3	
4	5	6	
7	0	8	
Replace zero row column
----Printing Matrix-----
1	0	3	
4	0	6	
0	0	0	

Process finished with exit code 0

*/
