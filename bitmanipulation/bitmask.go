package bitmanipulation

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
	F3
	F4
	F5
	F6
	F7
)
const (
	Empty Bits = 0
	Full  Bits = 255
)

// Set the 1 bits in flags to 1 in mask
// Set(00000000, 00000001) =>  00000001
// Set(00000001, 00000011) =>  00000011
func Set(mask, flag Bits) Bits {
	// bitwise union
	return mask | flag
}

// Return 1 for bits that are 1 in both mask and flag
// Has(11111111, 000)
// Check if mask contains flag
// Nothing contains 0 (empty bitmask)
// 11111111 contains all non-empty flags
// Has is built for flag, any match is a match, this is some flag if overcrowd flag
// use ContainsAll(mask, flags) for checking for composite
func Has(mask, flag Bits) bool {
	return mask&flag != 0
}

func ContainsAll(mask, flags Bits) bool {
	return mask&flags == flags
}

// complement of mask (~mask)
// ex: Complement(0b00000000) => 0b11111111
func Complement(mask Bits) Bits {
	return mask ^ 255
}

// Endianness worth exploring as well
/*
data := []uint8{0x12, 0x34, 0x56, 0x78}

	// Convert to int32 (big-endian)
	value := binary.BigEndian.Uint32(data)
	fmt.Println("Big-endian value:", value) // Output: Big-endian value: 305419896

	// Convert to int32 (little-endian)
	value = binary.LittleEndian.Uint32(data)
	fmt.Println("Little-endian value:", value)
*/
