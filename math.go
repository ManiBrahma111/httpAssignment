package main

func sub(a, b int) int {
	return a - b
}
func add(a, b int) int {
	return a + b
}
func multi(a, b int) int {
	return a * b
}
func div(a, b int) int {
	if b == 0 {
		return 0
	} else {
		return a / b
	}

}
