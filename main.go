package main

func toCover(a bool) int {
	if a {
		return 42
	}
	panic("oh no!")
}

func main() {
}
