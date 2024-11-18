package gobip39

import (
	"crypto/sha512"
	"fmt"
	"strings"

	"github.com/viwet/GoBIP39/words"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/text/unicode/norm"
)

const (
	// BIP-39 Password prefix
	PasswordPrefix = "mnemonic"
	// PBKDF2 Iterations count
	Iterations = 2048
	// PBKDF2 Derived key length
	KeyLength = 64
)

// ExtractSeed returns seed derived from mnemonic and password
func ExtractSeed(mnemonic []string, list words.List, password string) ([]byte, error) {
	if err := ValidateMnemonic(mnemonic, list); err != nil {
		return nil, fmt.Errorf("cannot extract seed: %w", err)
	}

	return pbkdf2.Key(
		[]byte(norm.NFKD.String(strings.Join(mnemonic, " "))),
		[]byte(norm.NFKD.String(PasswordPrefix+password)),
		Iterations,
		KeyLength,
		sha512.New,
	), nil
}
