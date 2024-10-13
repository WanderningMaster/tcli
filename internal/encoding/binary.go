package encoding

import (
	"bytes"
	"context"
	"encoding/gob"
)

type BinaryParser struct {
	ext    string
	buffer *bytes.Buffer
	enc    *gob.Encoder
	dec    *gob.Decoder
}

func NewBinaryParser() *BinaryParser {
	var buff bytes.Buffer
	return &BinaryParser{
		ext:    "bin",
		enc:    gob.NewEncoder(&buff),
		dec:    gob.NewDecoder(&buff),
		buffer: &buff,
	}
}

func (p *BinaryParser) Marshal(ctx context.Context, v any) ([]byte, error) {
	p.buffer.Reset()
	err := p.enc.Encode(v)

	return p.buffer.Bytes(), err
}

func (p *BinaryParser) Unmarshal(ctx context.Context, data []byte, v any) error {
	p.buffer.Reset()
	p.buffer.Write(data)

	return p.dec.Decode(v)
}

func (p *BinaryParser) Extension() string {
	return p.ext
}
