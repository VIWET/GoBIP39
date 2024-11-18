package words

var (
	chineseSimplified list

	// ChineseSimplified BIP-39 words
	ChineseSimplified List
)

func init() {
	chineseSimplified = newList("bip39/chinese_simplified.txt", 0xe3721bbf)
	ChineseSimplified = func() *list { return &chineseSimplified }
}
