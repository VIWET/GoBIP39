package words

var (
	korean list

	// Korean BIP-39 words
	Korean List
)

func init() {
	korean = newList("bip39/korean.txt", 0x4ef461eb)
	Korean = func() *list { return &korean }
}
