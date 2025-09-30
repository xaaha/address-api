package scripts

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// GenerateKey generates Onetime API Key
func GenerateKey() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	encodedKey := base64.URLEncoding.EncodeToString(key)
	fmt.Println("Key: ")
	fmt.Println(encodedKey)
}
