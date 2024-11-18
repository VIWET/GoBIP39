package words

var (
	english list

	// English BIP-39 words
	English List
)

func init() {
	english = newList("bip39/english.txt", 0xc1dbd296)
	English = func() *list { return &english }
}
