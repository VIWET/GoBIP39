package words

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"hash/crc32"
	"path"
	"strings"

	"golang.org/x/text/unicode/norm"
)

//go:embed bip39/*.txt
var bip39 embed.FS

// Number of words in BIP-39 list
const BIP39WordsCount = 1 << 11

// List is a list getter function
type List func() *list

// BIP-0039 words list container
type list struct {
	// Language of list
	language string
	// Word by index
	words [BIP39WordsCount]string
	// Index by word
	indices map[string]int
}

// Language of the list
func (l *list) Language() string {
	return l.language
}

// At returns a word at the index
func (l *list) At(index int) (string, error) {
	if index < 0 || index >= BIP39WordsCount {
		return "", fmt.Errorf("invalid index: index must be positive and less 2048, got: %d", index)
	}

	return l.words[index], nil
}

// IndexOf returns an index of the word
func (l *list) IndexOf(word string) (int, error) {
	index, ok := l.indices[abbreviate(word)]
	if !ok {
		return -1, fmt.Errorf("invalid word: %s list does not contain word %s", l.language, word)
	}

	return index, nil
}

// abbreviate returns first 4 unique in a list characters of the word
func abbreviate(word string) string {
	word = norm.NFKC.String(word)
	unique := []rune(word)[:4]

	return string(unique)
}

func newList(fileName string, checksum uint32) list {
	data, err := bip39.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if checksum != crc32.ChecksumIEEE(data) {
		panic(fmt.Sprintf("%s invalid checksum", fileName))
	}

	var (
		scanner = bufio.NewScanner(bytes.NewBuffer(data))
		index   = 0

		language = strings.Split(path.Base(fileName), ".")[0]

		words   [BIP39WordsCount]string
		indices = make(map[string]int, BIP39WordsCount)
	)

	for scanner.Scan() {
		word := scanner.Text()
		words[index] = word
		indices[abbreviate(word)] = index

		index++
	}

	return list{
		language: language,
		words:    words,
		indices:  indices,
	}
}
