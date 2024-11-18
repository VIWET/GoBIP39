package words

var (
	spanish list

	// Spanish BIP-39 words
	Spanish List
)

func init() {
	spanish = newList("bip39/spanish.txt", 0x266e4f3d)
	Spanish = func() *list { return &spanish }
}
