package words

var (
	japanese list

	// Japanese BIP-39 words
	Japanese List
)

func init() {
	japanese = newList("bip39/japanese.txt", 0xacc1419)
	Japanese = func() *list { return &japanese }
}
