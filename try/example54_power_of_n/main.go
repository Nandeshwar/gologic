package main

func pow(b, p int) int {
	if p == 0 {
		return 1
	}
	
	// with less call
	if p % 2 == 0 {
		return pow(b*b, p/2)
	}
	return b * pow(b, p-1)
}