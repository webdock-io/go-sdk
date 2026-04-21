package accountpublickeys

import (
	"context"
	"strconv"
)

type DeletePublicOptions struct {
	ID int64
}

func (w *AccountPublicKeys) Delete(ctx context.Context, options DeletePublicOptions) error {
	_, err := w.client.Do(ctx, "DELETE", "v1/account/publicKeys/"+strconv.FormatInt(options.ID, 10), nil, nil)
	return err
}
