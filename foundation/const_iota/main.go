package main

const (
	PI = 3.14159265359
	E  = 2.71828182846
)

const (
	Zero = iota
	One
	Two
	Three
)

const (
	MulOne = 1 << iota
	_
	MulFour
	_
	MulSixteen
)

func main() {
	println(PI, E)                       // 3.14159265359 2.71828182846
	println(Zero, One, Two, Three)       // 0 1 2 3
	println(MulOne, MulFour, MulSixteen) // 1 4 16
}
