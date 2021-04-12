package e2etest

import (
	"context"

	"github.com/cenkalti/backoff/v4"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/telegram/internal/helpers"
	"github.com/gotd/td/tg"
)

type waitInvoker struct {
	prev tg.Invoker
}

func retryFloodWait(ctx context.Context, cb func() error) error {
	return backoff.Retry(func() error {
		if err := cb(); err != nil {
			if ok, err := helpers.FloodWait(ctx, err); ok {
				return err
			}

			return backoff.Permanent(err)
		}

		return nil
	}, backoff.WithContext(backoff.NewExponentialBackOff(), ctx))
}

func (w waitInvoker) InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	return retryFloodWait(ctx, func() error {
		return w.prev.InvokeRaw(ctx, input, output)
	})
}
