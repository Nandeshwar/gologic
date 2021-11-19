package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrimeNumbers(t *testing.T) {
	Convey("Testing prime numbers", t, func() {
		Convey("succes testing prime numbers algorithm1", func() {
			data := []struct {
				begin  int
				end    int
				output []int
			}{
				{begin: 1, end: 10, output: []int{2, 3, 5, 7}},
			}

			for _, d := range data {
				So(primeNumberBetween(d.begin, d.end), ShouldResemble, d.output)
			}
		})

		Convey("succes testing prime numbers algorithm2", func() {
			data := []struct {
				begin  int
				end    int
				output []int
			}{
				{begin: 1, end: 10, output: []int{2, 3, 5, 7}},
			}

			for _, d := range data {
				So(primeNumberBetweenAlgorith2(d.begin, d.end), ShouldResemble, d.output)
			}
		})
	})
}

func BenchmarkTestPrimeNumbersAlgorithm1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		primeNumberBetween(1, 100)
	}
}

func BenchmarkTestPrimeNumbersAlgorithm2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		primeNumberBetweenAlgorith2(1, 100)
	}
}

/*
bash-3.2$ go test -benchmem -run=^$ -bench .
goos: darwin
goarch: amd64
pkg: gologic/try/example66_prime_number.go
BenchmarkTestPrimeNumbersAlgorithm1-16        2118 ns/op   504 B/op 6 allocs/op
BenchmarkTestPrimeNumbersAlgorithm2-16        442 ns/op    616 B/op 7 allocs/op
PASS
ok
bash-3.2$
*/
