package gobip39

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"math/big"
)

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
var ErrInvalidEntropyBitlen = errors.New("entropy bitlen must be one of 128, 160, 192, 224 or 256")

type (
	// Initial entropy
	Entropy []byte
	// Entropy with checksum
	ChecksumEntropy []byte
)

// NewEntropy generates new random entropy of bitlen
func NewEntropy(bitlen uint) (Entropy, error) {
	if !isValidBitlen(bitlen) {
		return nil, ErrInvalidEntropyBitlen
	}

	return Entropy(RandomBytes(BitlenToBytelen(bitlen))), nil
}

// isValidBitlen returns false if the bitlen is less than 128, greater than 256, or not divisible by 32
func isValidBitlen(bitlen uint) bool {
	return bitlen >= MinBitlen && bitlen <= MaxBitlen && bitlen%BitlenDivisor == 0
}

// Bitlen return a bitlen of entropy
func (e Entropy) Bitlen() uint {
	return BytelenToBitlen(uint(len(e)))
}

// AddChecksum returns entropy with checksum
func (e Entropy) AddChecksum() ChecksumEntropy {
	var (
		hash           = sha256.Sum256(e)
		checksumBitlen = e.Bitlen() / BitlenDivisor
		shift          = 256 - checksumBitlen
	)

	checksum := new(big.Int).Rsh(new(big.Int).SetBytes(hash[:]), shift)

	return PadLeftToBitlen(
		new(big.Int).Add(
			new(big.Int).Lsh(
				new(big.Int).SetBytes(e),
				checksumBitlen,
			),
			checksum,
		).Bytes(),
		e.Bitlen()+checksumBitlen,
	)
}

// IsValidChecksum validates entorpy checksum
func (e Entropy) IsValidChecksum(checksum []byte) bool {
	var (
		hash           = sha256.Sum256(e)
		checksumBitlen = e.Bitlen() / BitlenDivisor
		shift          = 256 - checksumBitlen
	)

	return bytes.Equal(
		new(big.Int).Rsh(
			new(big.Int).SetBytes(hash[:]),
			shift,
		).Bytes(),
		checksum,
	)
}

// Bitlen returns a bitlen of entorpy with checksum
func (e ChecksumEntropy) Bitlen() uint {
	bitlen := BytelenToBitlen(uint(len(e)) - 1)
	return bitlen + bitlen/BitlenDivisor
}

// EntropyBitlen returns a bitlen of entropy without checksum
func (e ChecksumEntropy) EntropyBitlen() uint {
	return BytelenToBitlen(uint(len(e)) - 1)
}

// RemoveChecksum returns initial entropy and checksum separately
func (e ChecksumEntropy) RemoveChecksum() (Entropy, []byte) {
	var (
		bitlen         = e.Bitlen()
		entropyBitlen  = e.EntropyBitlen()
		checksumBitlen = bitlen - entropyBitlen
	)

	checksumEntropy := new(big.Int).SetBytes(e)

	entropy := new(big.Int).Rsh(
		checksumEntropy,
		checksumBitlen,
	).Bytes()

	checksum := new(big.Int).And(
		checksumEntropy,
		new(big.Int).SetUint64(uint64(1<<checksumBitlen)-1),
	).Bytes()

	return PadLeftToBitlen(entropy, entropyBitlen), checksum
}
