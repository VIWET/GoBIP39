package gobip39

import "errors"

const (
	// Minimum allowed by BIP-39 entropy bitlen
	MinBitlen = 128
	// Maximum allowed by BIP-39 entropy bitlen
	MaxBitlen = 256
	// Divisor that used to get checksum bits of SHA-256 hash
	BitlenDivisor = 32

	// Bits are splitted into groups of 11 bits, which are represents an index of the word
	GroupBitlen = 11
)

// ErrInvalidEntropyBitlen
var ErrInvalidEntropyBitlen = errors.New("entropy bitlen must be one of [128, 160, 192, 224, 256]")

// NewEntropy generates new random entropy of bitlen
func NewEntropy(bitlen uint) ([]byte, error) {
	if !isValidBitlen(bitlen) {
		return nil, ErrInvalidEntropyBitlen
	}

	return RandomBytes(BitlenToBytelen(bitlen)), nil
}

// isValidBitlen returns false if the bitlen is less than 128, greater than 256, or not divisible by 32
func isValidBitlen(bitlen uint) bool {
	return bitlen >= MinBitlen && bitlen <= MaxBitlen && bitlen%BitlenDivisor == 0
}
