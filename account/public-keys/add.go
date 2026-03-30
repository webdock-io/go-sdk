package accountpublickeys

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreatePublicKeyOptions struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

func (w *AccountPublicKeys) Add(opts CreatePublicKeyOptions) (*PublicKey, error) {
	var publicKey PublicKey

	byts, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request body")
	}
	_, err = w.client.Do("POST", "v1/account/publicKeys", bytes.NewReader(byts), &publicKey)
	if err != nil {
		return nil, err
	}
	return &publicKey, nil
}
