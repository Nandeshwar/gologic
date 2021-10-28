package main

// https://www.youtube.com/watch?v=dzYq5VEMZIg
// n = 5, k =3
// 0 1 2 3 4
// 3rd item will be removed
// 3 4  + 0 1 = 3 4 0 1
// 1+ 3 4  = 1 3 4
// 1 3
// 3
func JosephProblem(n, k int) int {

	if n == 1 {
		return 0
	}

	x := JosephProblem(n-1, k)

	y := (x + k) % n

	return y

}
