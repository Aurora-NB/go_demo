package main

import (
	"fmt"
	"math"
)

func main() {

	var b1 bool = true
	var b2 bool
	fmt.Println(b1, b2)
	fmt.Println()

	var f32 float32 = 3.141592611111      // 小数点7位
	var f64 float64 = 3.12345678901234567 // 小数点15位
	fmt.Println(f32, f64)
	fmt.Printf("f32 -> %T , f64 > %T\n", f32, f64)
	fmt.Println()

	var by byte = 97
	fmt.Printf("%c , %d", by, by)
	fmt.Println()

	var i int = math.MaxInt
	var ui uint = math.MaxUint
	var i8 int8 = math.MaxInt8
	var ui8 uint8 = math.MaxUint8
	var i16 int16 = math.MaxInt16
	var ui16 uint16 = math.MaxUint16
	var i32 int32 = math.MaxInt32
	var ui32 uint32 = math.MaxUint32
	var i64 int64 = math.MaxInt64
	var ui64 uint64 = math.MaxUint64
	fmt.Println(i)
	fmt.Println(ui)
	fmt.Println(i8)
	fmt.Println(ui8)
	fmt.Println(i16)
	fmt.Println(ui16)
	fmt.Println(i32)
	fmt.Println(ui32)
	fmt.Println(i64)
	fmt.Println(ui64)

	var c64 complex64
	var c128 complex128
	fmt.Println(c64)
	fmt.Println(c128)
	fmt.Println(real(c128), imag(c128))
	fmt.Println()

	var str string = "hello world"
	fmt.Printf("str = %s , len = %d\n", str, len(str))
}
