package encoding

import (
	"context"
	"encoding/json"
)

type JsonParser struct {
	ext string
}

func NewJsonParser() *JsonParser {
	return &JsonParser{
		ext: "json",
	}
}

func (p *JsonParser) Marshal(ctx context.Context, v any) ([]byte, error) {
	return json.Marshal(v)
}

func (p *JsonParser) Unmarshal(ctx context.Context, data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (p *JsonParser) Extension() string {
	return p.ext
}
