package profiles

import (
	"context"
	"fmt"
)

type DeleteProfileOptions struct {
	ProfileSlug string
}

func (p *Profiles) Delete(ctx context.Context, opts DeleteProfileOptions) error {
	_, err := p.client.Do(ctx, "DELETE", fmt.Sprintf("/profiles/%s", opts.ProfileSlug), nil, nil)
	return err
}
