package gobip39

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/viwet/GoBIP39/words"
)

const (
	// Minimum allowed by BIP-39 mnemonic length
	MinLength = 12
	// Maximum allowed by BIP-39 mnemonic length
	MaxLength = 24

	// Mnemonic length divisor
	LengthDivisor = 3
)

// ErrInvalidMnemonicLength
var ErrInvalidMnemonicLength = errors.New("mnemonic length must be one of 12, 15, 18, 21 or 24")

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
	for index := int(length - 1); index >= 0; index-- {
		wordIndex := new(big.Int).And(e, mask)
		e.Rsh(e, GroupBitlen)

		// wordIndex is always positive integer less than 2048
		word, err := words.At(int(wordIndex.Uint64()))
		if err != nil {
			return nil, fmt.Errorf("cannot extract mnemonic: %w", err)
		}

		mnemonic[index] = word
	}

	return mnemonic, nil
}
