package pystruct

import "fmt"
import "testing"
import "reflect"

var bArrayBE []byte
var bArrayLE []byte

var tbool bool = true
var tchar int32 = 'c'
var tbyte byte = 1
var tubyte byte = 1
var tshort int16 = 2
var tushort uint16 = 2
var tint int32 = 4
var tuint uint32 = 4
var tlong int32 = 4
var tulong uint32 = 4
var tllong int64 = 8
var tullong uint64 = 8
var tfloat float32 = 4
var tdouble float64 = 8

func TestPack(t *testing.T) {
	// Table

	// ? (boolean)       = size 1 = bool or int, int16, int32, int64
	// c (char)          = size 1 = char or int
	// b (byte)          = size 1 = byte or uint
	// B (u. byte)       = size 1 = byte or uint
	// h (short)         = size 2 = int16
	// h (u. short)      = size 2 = uint16
	// i (int)	         = size 4 = int32
	// I (u. int)        = size 4 = uint32
	// l (long)          = size 4 = int32
	// L (u. long)       = size 4 = uint32
	// q (long long)     = size 8 = int64
	// Q (u. long long)  = size 8 = uint64
	// f (float)         = size 4 = float32
	// d (double)		 = size 8 = float32

	print("Packing values... ")

	bArrayLE = Pack("<?cbBhHiIlLqQfd", tbool, tchar, tbyte, tubyte, tshort, tushort, tint, tuint, tlong, tulong, tllong, tullong, tfloat, tdouble)
	bArrayBE = Pack("!?cbBhHiIlLqQfd", tbool, tchar, tbyte, tubyte, tshort, tushort, tint, tuint, tlong, tulong, tllong, tullong, tfloat, tdouble)

	println("OK")
}

func TestUnpack(t *testing.T) {
	println("Trying to unpack values Little Endian\n")
	if uLe, err := Unpack("<?cbBhHiIlLqQfd", bArrayLE); err == nil {

		Bool := reflect.ValueOf(uLe[0]).Bool()
		Char := int32(reflect.ValueOf(uLe[1]).Uint())
		Byte := byte(reflect.ValueOf(uLe[2]).Uint())
		UByte := byte(reflect.ValueOf(uLe[3]).Uint())
		Short := int16(reflect.ValueOf(uLe[4]).Int())
		UShort := uint16(reflect.ValueOf(uLe[5]).Uint())
		Int := int32(reflect.ValueOf(uLe[6]).Uint())
		UInt := uint32(reflect.ValueOf(uLe[7]).Uint())
		Long := int32(reflect.ValueOf(uLe[8]).Int())
		ULong := uint32(reflect.ValueOf(uLe[9]).Int())
		LLong := int64(reflect.ValueOf(uLe[10]).Int())
		ULLong := uint64(reflect.ValueOf(uLe[11]).Uint())
		Float := float32(reflect.ValueOf(uLe[12]).Float())
		Double := float64(reflect.ValueOf(uLe[13]).Float())

		fmt.Printf("%v == %v : %v\n", Bool, tbool, Bool == tbool)
		fmt.Printf("%v == %v : %v\n", string(Char), string(tchar), Char == tchar)
		fmt.Printf("%v == %v : %v\n", Byte, tbyte, Byte == tbyte)
		fmt.Printf("%v == %v : %v\n", UByte, tubyte, UByte == tubyte)
		fmt.Printf("%v == %v : %v\n", Short, tshort, Short == tshort)
		fmt.Printf("%v == %v : %v\n", UShort, tushort, UShort == tushort)
		fmt.Printf("%v == %v : %v\n", Int, tint, Int == tint)
		fmt.Printf("%v == %v : %v\n", UInt, tuint, UInt == tuint)
		fmt.Printf("%v == %v : %v\n", Long, tlong, Long == tlong)
		fmt.Printf("%v == %v : %v\n", ULong, tulong, ULong == tulong)
		fmt.Printf("%v == %v : %v\n", LLong, tllong, LLong == tllong)
		fmt.Printf("%v == %v : %v\n", ULLong, tullong, ULLong == tullong)
		fmt.Printf("%.2f == %.2f : %v\n", Float, tfloat, Float == tfloat)
		fmt.Printf("%.2f == %.2f : %v\n", Double, tdouble, Double == tdouble)

	} else {
		t.Fatalf("%s", err.Error())
	}

	println("\nLittle Endian OK\n\nTrying to unpack values Big Endian...\n")
	if uBe, err := Unpack("!?cbBhHiIlLqQfd", bArrayBE); err == nil {

		Bool := reflect.ValueOf(uBe[0]).Bool()
		Char := int32(reflect.ValueOf(uBe[1]).Uint())
		Byte := byte(reflect.ValueOf(uBe[2]).Uint())
		UByte := byte(reflect.ValueOf(uBe[3]).Uint())
		Short := int16(reflect.ValueOf(uBe[4]).Int())
		UShort := uint16(reflect.ValueOf(uBe[5]).Uint())
		Int := int32(reflect.ValueOf(uBe[6]).Uint())
		UInt := uint32(reflect.ValueOf(uBe[7]).Uint())
		Long := int32(reflect.ValueOf(uBe[8]).Int())
		ULong := uint32(reflect.ValueOf(uBe[9]).Int())
		LLong := int64(reflect.ValueOf(uBe[10]).Int())
		ULLong := uint64(reflect.ValueOf(uBe[11]).Uint())
		Float := float32(reflect.ValueOf(uBe[12]).Float())
		Double := float64(reflect.ValueOf(uBe[13]).Float())

		fmt.Printf("%v == %v : %v\n", Bool, tbool, Bool == tbool)
		fmt.Printf("%v == %v : %v\n", string(Char), string(tchar), Char == tchar)
		fmt.Printf("%v == %v : %v\n", Byte, tbyte, Byte == tbyte)
		fmt.Printf("%v == %v : %v\n", UByte, tubyte, UByte == tubyte)
		fmt.Printf("%v == %v : %v\n", Short, tshort, Short == tshort)
		fmt.Printf("%v == %v : %v\n", UShort, tushort, UShort == tushort)
		fmt.Printf("%v == %v : %v\n", Int, tint, Int == tint)
		fmt.Printf("%v == %v : %v\n", UInt, tuint, UInt == tuint)
		fmt.Printf("%v == %v : %v\n", Long, tlong, Long == tlong)
		fmt.Printf("%v == %v : %v\n", ULong, tulong, ULong == tulong)
		fmt.Printf("%v == %v : %v\n", LLong, tllong, LLong == tllong)
		fmt.Printf("%v == %v : %v\n", ULLong, tullong, ULLong == tullong)
		fmt.Printf("%.2f == %.2f : %v\n", Float, tfloat, Float == tfloat)
		fmt.Printf("%.2f == %.2f : %v\n", Double, tdouble, Double == tdouble)

	} else {
		t.Fatalf("%s", err.Error())
	}

	println("\nBig Endian OK")
}