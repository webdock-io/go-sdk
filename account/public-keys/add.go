package accountpublickeys

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

type CreatePublicKeyOptions struct {
	Name      string `json:"name"`
	PublicKey string `json:"publicKey"`
}

func (w *AccountPublicKeys) Add(ctx context.Context, opts CreatePublicKeyOptions) (*PublicKey, error) {
	var publicKey PublicKey

	byts, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the request body")
	}
	_, err = w.client.Do(ctx, "POST", "v1/account/publicKeys", bytes.NewReader(byts), &publicKey)
	if err != nil {
		return nil, err
	}
	return &publicKey, nil
}
