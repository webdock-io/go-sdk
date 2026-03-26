package accountpublickeys

import (
	"strconv"

	"github.com/webdock-io/go-sdk/client"
)

type DeletePublicOptions struct {
	ID int64
}

func (w *AccountPublicKeys) Delete(options DeletePublicOptions) error {
	_, err := client.Do("DELETE", "v1/account/publicKeys/"+strconv.FormatInt(options.ID, 10), nil, nil)
	return err
}
