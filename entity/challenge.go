package entity

import (
	"encoding/json"
)

// Challenge represents challenge generated from server to client to perform the PoW
type Challenge struct {
	Challenge      []byte `json:"challenge"`
	Difficulty     int    `json:"difficulty"`
	SolutionLength int    `json:"solution_length"`
}

// EncodeToString encodes challenge to string
func (challenge Challenge) EncodeToString() (string, error) {
	challangeBytes, err := json.Marshal(challenge)
	if err != nil {
		return "", err
	}

	return string(challangeBytes), nil
}

// DecodeFromBytes decodes challenge from bytes
func (challenge *Challenge) DecodeFromBytes(data []byte) (err error) {
	err = json.Unmarshal(data, challenge)
	if err != nil {
		return
	}
	return
}
