package main

func main() {
	println("Hello, World!")
}

type ID interface {
	IsUnique() bool
}
