package gobip39_test

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"os"
	"path"
	"slices"
	"strings"
	"testing"

	bip39 "github.com/viwet/GoBIP39"
	"github.com/viwet/GoBIP39/words"
	"golang.org/x/text/unicode/norm"
)

func Test_BIP39(t *testing.T) {
	f := func(t *testing.T, fileName string, list words.List) {
		tests := LoadTestVector(t, fileName)
		name := strings.Split(path.Base(fileName), ".")[0]
		for _, test := range tests {
			t.Run(name, func(t *testing.T) {
				mnemonic, err := bip39.ExtractMnemonic(test.Entropy, list)
				if err != nil {
					t.Fatal(err)
				}

				if !slices.Equal(test.Mnemonic, mnemonic) {
					t.Fatalf(
						"invalid mnemonic extracted:\nWant: %s\nGot:  %s",
						strings.Join(test.Mnemonic, " "),
						strings.Join(mnemonic, " "),
					)
				}

				entropy, err := bip39.ExtractEntropy(test.Mnemonic, list)
				if err != nil {
					t.Fatal(err)
				}

				if !bytes.Equal(test.Entropy, entropy) {
					t.Fatalf(
						"invalid entropy extracted:\nWant: %x\nGot:  %x",
						test.Entropy,
						entropy,
					)
				}

				seed, err := bip39.ExtractSeed(test.Mnemonic, list, "TREZOR")
				if err != nil {
					t.Fatal(err)
				}

				if !bytes.Equal(test.Seed, seed) {
					t.Fatalf(
						"invalid seed extracted:\nWant: %x\nGot:  %x",
						test.Seed,
						seed,
					)
				}
			})
		}
	}

	f(t, "tests/chinese_simplified.json", words.ChineseSimplified)
	f(t, "tests/chinese_traditional.json", words.ChineseTraditional)
	f(t, "tests/czech.json", words.Czech)
	f(t, "tests/english.json", words.English)
	f(t, "tests/french.json", words.French)
	f(t, "tests/italian.json", words.Italian)
	f(t, "tests/japanese.json", words.Japanese)
	f(t, "tests/korean.json", words.Korean)
	f(t, "tests/portuguese.json", words.Portuguese)
	f(t, "tests/spanish.json", words.Spanish)
}

type Test struct {
	Entropy  []byte
	Mnemonic []string
	Seed     []byte
}

func LoadTestVector(t *testing.T, path string) []Test {
	t.Helper()

	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	var vector [][]string
	if err := json.NewDecoder(file).Decode(&vector); err != nil {
		t.Fatal(err)
	}

	tests := make([]Test, len(vector))
	for i, test := range vector {
		entropy, err := hex.DecodeString(test[0])
		if err != nil {
			t.Fatal(err)
		}

		mnemonic := bip39.NormalizeMnemonic(strings.Split(norm.NFKD.String(test[1]), " "))

		seed, err := hex.DecodeString(test[2])
		if err != nil {
			t.Fatal(err)
		}

		tests[i] = Test{
			Mnemonic: mnemonic,
			Entropy:  entropy,
			Seed:     seed,
		}
	}

	return tests
}
