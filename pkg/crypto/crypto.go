package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func GetRandomBytes(len int) []byte {
	bytes := make([]byte, len)

	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatalf("error while generating random string: %s", err)
	}

	return bytes
}

// Function to generate the hash of a given data string
func CalculateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
