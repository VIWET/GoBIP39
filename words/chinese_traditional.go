package words

var (
	chineseTraditional list

	// ChineseTraditional BIP-39 words
	ChineseTraditional List
)

func init() {
	chineseTraditional = newList("bip39/chinese_traditional.txt", 0x3c20b443)
	ChineseTraditional = func() *list { return &chineseTraditional }
}
