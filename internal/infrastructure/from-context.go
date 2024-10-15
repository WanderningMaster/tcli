package infrastructure

import "context"

type storageKey struct{}

func WithStorage(ctx context.Context, s *Storage) context.Context {
	return context.WithValue(ctx, storageKey{}, s)
}

func GetStorage(ctx context.Context) *Storage {
	logger := ctx.Value(storageKey{})

	return logger.(*Storage)
}
