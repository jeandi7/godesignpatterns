package main

func main() {
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addN(5)
	result := addFive(7)
	println(result)

}
