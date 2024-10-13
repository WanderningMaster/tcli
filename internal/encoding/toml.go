package encoding

import (
	"context"

	"github.com/BurntSushi/toml"
)

type TomlParser struct {
	ext string
}

func NewTomlParser() *TomlParser {
	return &TomlParser{
		ext: "toml",
	}
}

func (p *TomlParser) Marshal(ctx context.Context, v any) ([]byte, error) {
	return toml.Marshal(v)
}

func (p *TomlParser) Unmarshal(ctx context.Context, data []byte, v any) error {
	return toml.Unmarshal(data, v)
}

func (p *TomlParser) Extension() string {
	return p.ext
}
