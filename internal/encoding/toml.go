package encoding

import (
	"context"

	"github.com/BurntSushi/toml"
)

type TomlParser struct {
}

func NewTomlParser() *TomlParser {
	return &TomlParser{}
}

func (p *TomlParser) Marshal(ctx context.Context, v any) ([]byte, error) {
	return toml.Marshal(v)
}

func (p *TomlParser) Unmarshal(ctx context.Context, data []byte, v any) error {
	return toml.Unmarshal(data, v)
}
