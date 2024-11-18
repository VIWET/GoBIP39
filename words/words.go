package words

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

// Number of words in BIP-39 list
const BIP39WordsCount = 1 << 11

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
