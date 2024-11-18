package words

var (
	english list
	English List
)

func init() {
	english = newList("bip39/english.txt", 0xc1dbd296)
	English = func() *list { return &english }
}
