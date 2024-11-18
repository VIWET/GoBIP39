package words

var (
	portuguese list

	// Portuguese BIP-39 words
	Portuguese List
)

func init() {
	portuguese = newList("bip39/portuguese.txt", 0xe627a546)
	Portuguese = func() *list { return &portuguese }
}
