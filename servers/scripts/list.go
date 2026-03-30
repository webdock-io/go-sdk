package serverscripts

import "fmt"

type ListScriptsOptions struct {
	ServerSlug string
}

func (s *ServerScripts) List(opts ListScriptsOptions) ([]ServerScriptDTO, error) {
	var out []ServerScriptDTO
	_, err := s.client.Do("GET", fmt.Sprintf("v1/servers/%s/scripts", opts.ServerSlug), nil, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
