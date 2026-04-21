package accountpublickeys

import "context"

type PublicKey struct {
	ID      int64  `json:"id" tfsdk:"id"`
	Name    string `json:"name" tfsdk:"name"`
	Key     string `json:"key" tfsdk:"key"`
	Created string `json:"created" tfsdk:"created"`
}

func (w *AccountPublicKeys) List(ctx context.Context) (*[]PublicKey, error) {
	var out []PublicKey
	_, err := w.client.Do(ctx, "GET", "v1/account/publicKeys", nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
