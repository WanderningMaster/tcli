package encoding

import "context"

type Parser interface {
	Marshal(ctx context.Context, v any) ([]byte, error)
	Unmarshal(ctx context.Context, data []byte, v any) error
	Extension() string
}
