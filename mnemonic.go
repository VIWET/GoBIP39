package gobip39

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/viwet/GoBIP39/words"
	"golang.org/x/text/unicode/norm"
)

const (
	// Minimum allowed by BIP-39 mnemonic length
	MinLength = 12
	// Maximum allowed by BIP-39 mnemonic length
	MaxLength = 24

	// Mnemonic length divisor
	LengthDivisor = 3
)

var (
	// ErrInvalidMnemonicLength
	ErrInvalidMnemonicLength = errors.New("mnemonic length must be one of 12, 15, 18, 21 or 24")
	// ErrInvalidMnemonicChecksum
	ErrInvalidMnemonicChecksum = errors.New("invalid mnemonic checksum")
)

// ExtractMnemonic returns mnemonic phrase based on given entropy and language
func ExtractMnemonic(entropy Entropy, list words.List) ([]string, error) {
	var (
		checksumEntropy = entropy.AddChecksum()

		length   = checksumEntropy.Bitlen() / GroupBitlen
		mnemonic = make([]string, length)
		words    = list()

		e    = new(big.Int).SetBytes(checksumEntropy)
		mask = new(big.Int).SetUint64(uint64(1<<GroupBitlen) - 1)
	)

	// length is always positive integer in range between 12 and 24
	// #nosec
	for index := int(length - 1); index >= 0; index-- {
		wordIndex := new(big.Int).And(e, mask)
		e.Rsh(e, GroupBitlen)

		// wordIndex is always positive integer less than 2048
		word, err := words.At(int(wordIndex.Uint64())) // #nosec
		if err != nil {
			return nil, fmt.Errorf("cannot extract mnemonic: %w", err)
		}

		mnemonic[index] = word
	}

	return mnemonic, nil
}

// ExtractEntropy returns entropy based on given mnemonic phrase and language
func ExtractEntropy(mnemonic []string, list words.List) (Entropy, error) {
	if !isValidMnemonicLength(len(mnemonic)) {
		return nil, ErrInvalidMnemonicLength
	}

	var (
		words  = list()
		bitlen = uint(len(mnemonic)) * GroupBitlen
		e      = new(big.Int).SetUint64(0)
	)

	for _, word := range mnemonic {
		wordIndex, err := words.IndexOf(word)
		if err != nil {
			return nil, fmt.Errorf("cannot extract entropy: %w", err)
		}
		e.Lsh(e, GroupBitlen)
		// wordIndex is always positive integer less than 2048
		e.Add(e, new(big.Int).SetUint64(uint64(wordIndex))) // #nosec
	}

	checksumEntorpy := ChecksumEntropy(PadLeftToBitlen(e.Bytes(), bitlen))
	entropy, checksum := checksumEntorpy.RemoveChecksum()

	if !entropy.IsValidChecksum(checksum) {
		return nil, ErrInvalidMnemonicChecksum
	}

	return entropy, nil
}

// NormalizeMnemonic returns NFKD normalized UTF-8 encoded mnemonic
func NormalizeMnemonic(mnemonic []string) []string {
	for i, word := range mnemonic {
		mnemonic[i] = norm.NFKD.String(word)
	}

	return mnemonic
}

// ValidateMnemonic returns error if mnemonic has invalid length or it is impossible to extract entropy
func ValidateMnemonic(mnemonic []string, list words.List) error {
	_, err := ExtractEntropy(NormalizeMnemonic(mnemonic), list)
	if err != nil {
		return fmt.Errorf("mnemonic validation error: %w", err)
	}

	return nil
}

// isValidMnemonicLength returns false if mnemonic length is less than 12, greater than 24, or not divisible by 3
func isValidMnemonicLength(length int) bool {
	return length >= MinLength && length <= MaxLength && length%LengthDivisor == 0
}
