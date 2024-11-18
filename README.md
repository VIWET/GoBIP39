# Go BIP39

Golang implementation of BIP-0039 deterministic mnemonic code and key generation algorithm

# Example

Genetate new mnemonic and extract seed

```Go
package main

import (
	"log"

	bip39 "github.com/viwet/GoBIP39"
	"github.com/viwet/GoBIP39/words"
)

const Password = ""

func main() {
	entorpy, err := bip39.NewEntropy(256)
	if err != nil {
		log.Fatal(err)
	}

	mnemonic, err := bip39.ExtractMnemonic(entorpy, words.English)
	if err != nil {
		log.Fatal(err)
	}

	seed, err := bip39.ExtractSeed(mnemonic, words.English, Password)
	if err != nil {
		log.Fatal(err)
	}

	...
}
```

Extract seed from existing mnemonic

```Go
package main

import (
	"log"

	bip39 "github.com/viwet/GoBIP39"
	"github.com/viwet/GoBIP39/words"
)

const (
    Mnemonic = "あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あいこくしん　あおぞら"
    Password = "password"
)

func main() {
	mnemonic := bip39.SplitMnemonic(Mnemonic)

	seed, err := bip39.ExtractSeed(mnemonic, words.Japanse, Password)
	if err != nil {
		log.Fatal(err)
	}

	...
}
```
