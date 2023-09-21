package pow

import (
	"math/rand"
	"strings"
	"test-faraway/pkg/crypto"
	"time"
)

// GenerateChallengeStr generates a random challenge with the given min and max lengths
func GenerateChallengeStr(minLength, maxLength int) []byte {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	challegeLen := rand.Intn(maxLength-minLength) + minLength // len = [10, 20]

	return crypto.GetRandomBytes(challegeLen)
}

// Function to perform Proof-of-Work
func PerformPoW(challenge []byte, difficulty, solutionLength int) []byte {
	var solution = make([]byte, solutionLength)
	for {
		// Append a nonce value to the data string
		solution = crypto.GetRandomBytes(solutionLength)

		if VerifySolution(challenge, solution, difficulty) {
			return solution
		}
	}
}

func VerifySolution(challenge, solution []byte, difficulty int) bool {

	targetPrefix := strings.Repeat("0", difficulty) // target prefix containing the required number of leading zeros

	hash := crypto.CalculateHash(string(challenge) + string(solution))

	return strings.HasPrefix(hash, targetPrefix)
}
