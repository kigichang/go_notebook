package main

// Declaration of variables

func open() (uintptr, error) {
	return 0, nil
}

func main() {

	var a int
	var b int = 10
	var c = 20
	d := 30

	println(a, b, c, d) // 0 10 20 30

	var ptr, err = open()
	println(ptr, err) // 0x0, 0x0

	var m = map[string]int{"one": 1, "two": 2}

	val, ok := m["one"]
	println(val, ok) // 1, true

	val, ok = m["three"]
	println(val, ok) // 0, false

	x := 1 + 2i      // complex128
	y := 3 + 4i      // complex128
	z := x + y       // complex128
	println(x, y, z) // (+1.000000e+000+2.000000e+000i) (+3.000000e+000+4.000000e+000i) (+4.000000e+000+6.000000e+000i)

	x = complex(1, 2)
	y = complex(3, 4)
	z = x * y
	println(z, real(z), imag(z))
}
