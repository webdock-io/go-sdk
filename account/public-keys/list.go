package accountpublickeys

import (
	"github.com/webdock-io/go-sdk/client"
)

type PublicKey struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Key     string `json:"key"`
	Created string `json:"created"`
}

func (w *AccountPublicKeys) List() (*[]PublicKey, error) {
	var out []PublicKey
	_, err := client.Do("GET", "v1/account/publicKeys", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
