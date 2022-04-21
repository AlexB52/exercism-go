package variablelengthquantity

func EncodeVarint(input []uint32) []byte {
	var result []byte
	for _, n := range input {
		number := []byte{byte(n % 128)}
		n /= 128
		for n > 0 {
			number = append(number, byte(n%128+128))
			n /= 128
		}

		for i, j := 0, len(number)-1; i < j; i, j = i+1, j-1 {
			number[i], number[j] = number[j], number[i]
		}

		result = append(result, number...)
	}
	return result
}

func DecodeVarint(input []byte) ([]uint32, error) {
	// fmt.Println(input)
	// // a := input[0]
	// aString := fmt.Sprintf("%08b", 128)
	// // c := fmt.Sprintf("%x", a)
	// fmt.Println(byte(255))
	// fmt.Println("binary", aString)
	// b, _ := strconv.ParseInt(aString, 2, 16)
	// fmt.Println("hex", b)
	// fmt.Println("byte", byte(b))

	return nil, nil
}

// fmt.Println("uint 0x40", uint(0x40))
// fmt.Println("uint 0x00", uint(0x00))
// fmt.Println("uint 0xff", uint(0xff))
// fmt.Println("byte 0x40", byte(0x40))
// fmt.Println("byte 0x00", byte(0x00))
// fmt.Println("byte 0xff", uint32(0xfffFFFFF))
// fmt.Println("byte 0x81", int(0xff))
// fmt.Println("byte 0xff", math.MaxUint32)

// rawHex := "60A100"
// i, err := strconv.ParseUint(rawHex, 16, 32)
// if err != nil {
// 	fmt.Printf("%s", err)
// }
// fmt.Printf("%024b\n", i)
// fmt.Println("i", i)

// i2, err := strconv.ParseUint("40", 16, 32)
// if err != nil {
// 	fmt.Printf("%s", err)
// }
// fmt.Printf("%024b\n", i2)
// fmt.Println("i2", i2)
