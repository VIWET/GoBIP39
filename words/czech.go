package words

var (
	czech list

	// Czech BIP-39 words
	Czech List
)

func init() {
	czech = newList("bip39/czech.txt", 0xd1b5fda0)
	Czech = func() *list { return &czech }
}
