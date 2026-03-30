package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CreateEventHookOptions struct {
	CallbackUrl string  `json:"callbackUrl"`
	CallbackId  *string `json:"callbackId,omitempty"`
	EventType   *string `json:"eventType,omitempty"`
}

func (s *Hooks) Create(opts CreateEventHookOptions) (*EventHookDTO, error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}
	var out EventHookDTO
	_, err = s.client.Do("POST", "v1/hooks", bytes.NewBuffer(data), &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
