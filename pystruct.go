package pystruct

// TABLE
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

import "fmt"
import "math"
import "bytes"
import "errors"
import "reflect"
import "encoding/binary"

var size map[byte]int = map[byte]int{
	'?': 1, 'c': 1, 'b': 1, 'B': 1, 'h': 2,
	'H': 2, 'i': 4, 'I': 4, 'l': 4, 'L': 4,
	'q': 8, 'Q': 8, 'f': 4, 'd': 8,
}

func Pack(format string, values ...interface{}) []byte {

	var endian binary.ByteOrder = binary.LittleEndian

	if format[0] == '<' || format[0] == '>' || format[0] == '!' {

		if format[0] != '<' {
			endian = binary.BigEndian
		}

		format = format[1:len(format)]
	}

	buffer := new(bytes.Buffer)
	for i := range values {
		v := reflect.ValueOf(values[i])

		switch format[i] {
		case 's', 'S':
			binary.Write(buffer, endian, []byte(v.String()))
			break
		case 'c', 'b', 'B':
			if format[i] == 'c' {
				binary.Write(buffer, endian, byte(v.Int()))
			} else {
				binary.Write(buffer, endian, byte(v.Uint()))
			}
			break
		case 'h', 'H':
			if format[i] == 'h' {
				binary.Write(buffer, endian, int16(v.Int()))
			} else {
				binary.Write(buffer, endian, uint16(v.Uint()))
			}
			break
		case 'i', 'I', 'l', 'L':
			if format[i] == 'i' || format[i] == 'l' {
				binary.Write(buffer, endian, int32(v.Int()))
			} else {
				binary.Write(buffer, endian, uint32(v.Uint()))
			}
			break
		case 'q', 'Q':
			if format[i] == 'q' {
				binary.Write(buffer, endian, int64(v.Int()))
			} else {
				binary.Write(buffer, endian, uint64(v.Uint()))
			}
			break
		case 'f', 'd':
			if format[i] == 'f' {
				binary.Write(buffer, endian, float32(v.Float()))
			} else {
				binary.Write(buffer, endian, float64(v.Float()))
			}
		case '?':
			n := 0
			switch v.Type().Kind() {
			case reflect.Bool:
				if v.Bool() {
					n = 1
				}
				break
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				if v.Int() > 0 {
					n = 1
				}
				break
			}
			binary.Write(buffer, endian, byte(n))
			break
		}
	}
	return buffer.Bytes()
}

func Unpack(format string, buffer []byte) ([]interface{}, error) {

	var endian binary.ByteOrder = binary.LittleEndian

	if format[0] == '<' || format[0] == '>' || format[0] == '!' {

		if format[0] != '<' {
			endian = binary.BigEndian
		}

		format = format[1:len(format)]
	}

	s := 0
	for c := range format {
		s += size[format[c]]
	}

	buf := []interface{}{}

	if s == len(buffer) {

		//str := ""

		for c := range format {
			char := format[c]

			/*if char != 's' && len(str) > 0 {
				buf = append(buf, str)
				str = ""
			}*/

			switch char {
			//case 's':
			//	str += string(buffer[0])
			//	break
			case '?':
				b := byte(buffer[0])
				if b == 0 {
					buf = append(buf, false)
				} else {
					buf = append(buf, true)
				}
				break
			case 'c', 'b', 'B':
				buf = append(buf, byte(buffer[0]))
				break
			case 'h', 'H':
				b := buffer[0:size[char]]
				if char == 'h' {
					buf = append(buf, int16(endian.Uint16(b)))
				} else {
					buf = append(buf, endian.Uint16(b))
				}
				break
			case 'i', 'I', 'l', 'L':
				b := buffer[0:size[char]]
				if char == 'l' || char == 'L' {
					buf = append(buf, int32(endian.Uint32(b)))
				} else {
					buf = append(buf, endian.Uint32(b))
				}
				break
			case 'q', 'Q':
				b := buffer[0:size[char]]
				if char == 'q' {
					buf = append(buf, int64(endian.Uint64(b)))
				} else {
					buf = append(buf, endian.Uint64(b))
				}
				break
			case 'f', 'd':
				b := buffer[0:size[char]]
				if char == 'f' {
					buf = append(buf, math.Float32frombits(endian.Uint32(b)))
				} else {
					buf = append(buf, math.Float64frombits(endian.Uint64(b)))
				}
				break
			}
			buffer = buffer[size[char]:len(buffer)]
		}

		//if len(str) > 0 {
		//	buf = append(buf, str)
		//}

		return buf, nil
	} else {
		return buf, errors.New(fmt.Sprintf("pystruct: Unpack requires a buffer of length %d", s))
	}
}