package words

var (
	italian list

	// Italian BIP-39 words
	Italian List
)

func init() {
	italian = newList("bip39/italian.txt", 0x2fc7d07e)
	Italian = func() *list { return &italian }
}
