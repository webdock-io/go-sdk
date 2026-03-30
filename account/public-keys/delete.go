package accountpublickeys

import "strconv"

type DeletePublicOptions struct {
	ID int64
}

func (w *AccountPublicKeys) Delete(options DeletePublicOptions) error {
	_, err := w.client.Do("DELETE", "v1/account/publicKeys/"+strconv.FormatInt(options.ID, 10), nil, nil)
	return err
}
