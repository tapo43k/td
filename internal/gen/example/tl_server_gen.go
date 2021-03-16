// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

type ServerDispatcher struct {
	fallback func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error)
	handlers map[uint32]func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error)
}

func NewServerDispatcher(fallback func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error)) *ServerDispatcher {
	return &ServerDispatcher{
		fallback: fallback,
		handlers: map[uint32]func(context.Context, *bin.Buffer) (bin.Encoder, error){},
	}
}

func (s *ServerDispatcher) Handle(ctx context.Context, b *bin.Buffer) (bin.Encoder, error) {
	id, err := b.PeekID()
	if err != nil {
		return nil, err
	}

	f, ok := s.handlers[id]
	if !ok {
		return s.fallback(ctx, b)
	}

	return f(ctx, b)
}

func (s *ServerDispatcher) OnPing(f func(ctx context.Context, id int32) error) {
	handler := func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error) {
		var request PingRequest
		if err := request.Decode(b); err != nil {
			return nil, err
		}

		if err := f(ctx, request.ID); err != nil {
			return nil, err
		}

		return &Ok{}, nil
	}

	s.handlers[PingRequestTypeID] = handler
}

func (s *ServerDispatcher) OnSend(f func(ctx context.Context, msg SMS) (*SMS, error)) {
	handler := func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error) {
		var request SendRequest
		if err := request.Decode(b); err != nil {
			return nil, err
		}

		response, err := f(ctx, request.Msg)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	s.handlers[SendRequestTypeID] = handler
}

func (s *ServerDispatcher) OnSendMultipleSMS(f func(ctx context.Context, messages []SMS) error) {
	handler := func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error) {
		var request SendMultipleSMSRequest
		if err := request.Decode(b); err != nil {
			return nil, err
		}

		if err := f(ctx, request.Messages); err != nil {
			return nil, err
		}

		return &Ok{}, nil
	}

	s.handlers[SendMultipleSMSRequestTypeID] = handler
}

func (s *ServerDispatcher) OnDoAuth(f func(ctx context.Context) (AuthClass, error)) {
	handler := func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error) {
		var request DoAuthRequest
		if err := request.Decode(b); err != nil {
			return nil, err
		}

		response, err := f(ctx)
		if err != nil {
			return nil, err
		}
		return &AuthBox{Auth: response}, nil
	}

	s.handlers[DoAuthRequestTypeID] = handler
}

func (s *ServerDispatcher) OnEchoVector(f func(ctx context.Context, ids []int) ([]int, error)) {
	handler := func(ctx context.Context, b *bin.Buffer) (bin.Encoder, error) {
		var request EchoVectorRequest
		if err := request.Decode(b); err != nil {
			return nil, err
		}

		response, err := f(ctx, request.IDs)
		if err != nil {
			return nil, err
		}
		return &IntVector{Elems: response}, nil
	}

	s.handlers[EchoVectorRequestTypeID] = handler
}