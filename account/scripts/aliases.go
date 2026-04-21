package scripts

import "context"

type GetByIDOptions struct {
	ScriptID int64
}

type UpdateOptions struct {
	ScriptID int64
	Name     string
	Filename string
	Content  string
}

type DeleteOptions struct {
	ID int64
}

func (s *AccountScripts) GetByID(ctx context.Context, opts GetByIDOptions) (*AccountScriptDTO, error) {
	return s.GetAccountScriptById(ctx, GetAccountScriptByIdOptions{ScriptID: opts.ScriptID})
}

func (s *AccountScripts) Update(ctx context.Context, opts UpdateOptions) (*AccountScriptDTO, error) {
	return s.UpdateAccountScript(ctx, UpdateAccountScriptOptions{
		ScriptId: opts.ScriptID,
		Name:     opts.Name,
		Filename: opts.Filename,
		Content:  opts.Content,
	})
}

func (s *AccountScripts) Delete(ctx context.Context, opts DeleteOptions) error {
	return s.DeleteAccountScript(ctx, DeleteAccountScriptOptions{ScriptID: opts.ID})
}
