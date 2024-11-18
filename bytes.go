package gobip39

import "crypto/rand"

// RandomBytse generates random bytes sequence of given length
func RandomBytes(length uint) []byte {
	random := make([]byte, length)
	if _, err := rand.Read(random); err != nil {
		panic(err)
	}
	return random
}

// BitlenToBytelen converts a bitlet to a bytelen
func BitlenToBytelen(bitlen uint) uint {
	return (bitlen + 7) / 8
}

// BytelenToBitlen converts a bytelen to a bitlen
func BytelenToBitlen(bytelen uint) uint {
	return bytelen * 8
}

// PadLeftToBitlen pads a byte slice to a given bitlen if needed
func PadLeftToBitlen(bytes []byte, bitlen uint) []byte {
	var (
		bytelen = BitlenToBytelen(bitlen)
		offset  = bytelen - uint(len(bytes))
	)

	if offset <= 0 {
		return bytes
	}

	newBytes := make([]byte, bytelen)
	copy(newBytes[offset:], bytes)

	return newBytes
}
