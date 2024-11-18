package main

func main() {
	var a int
	var b uint
	var c int8
	var d uint8
	var e int16
	var f uint16
	var g int32
	var h uint32
	var i int64
	var j uint64
	var k float32
	var l float64
	var m complex64
	var n complex128

	var s string
	var p *int
	var sl []int
	var m1 map[string]int
	var ch chan int
	var f1 func(int) int
	var i1 interface{}

	println(a)         // 0
	println(b)         // 0
	println(c)         // 0
	println(d)         // 0
	println(e)         // 0
	println(f)         // 0
	println(g)         // 0
	println(h)         // 0
	println(i)         // 0
	println(j)         // 0
	println(k)         // 0
	println(l)         // 0
	println(m)         // (0+0i)
	println(n)         // (0+0i)
	println(s)         // ""
	println(p == nil)  // <nil>
	println(sl == nil) // []
	println(m1)        // map[]
	println(ch)        // <nil>
	println(f1)        // <nil>
	println(i1)        // <nil>

}
