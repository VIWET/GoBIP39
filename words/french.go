package words

var (
	french list

	// French BIP-39 words
	French List
)

func init() {
	french = newList("bip39/french.txt", 0x3e56b216)
	French = func() *list { return &french }
}
