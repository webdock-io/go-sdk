package serverscripts

import "fmt"

type GetScriptByIDOptions struct {
	ServerSlug string
	ScriptId   int
}

func (s *ServerScripts) GetByID(opts GetScriptByIDOptions) (*Script, error) {
	var out Script
	_, err := s.client.Do("GET", fmt.Sprintf("v1/servers/%s/scripts/%d", opts.ServerSlug, opts.ScriptId), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
