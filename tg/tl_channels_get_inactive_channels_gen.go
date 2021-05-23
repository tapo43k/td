// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

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
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// ChannelsGetInactiveChannelsRequest represents TL type `channels.getInactiveChannels#11e831ee`.
// Get inactive channels and supergroups
//
// See https://core.telegram.org/method/channels.getInactiveChannels for reference.
type ChannelsGetInactiveChannelsRequest struct {
}

// ChannelsGetInactiveChannelsRequestTypeID is TL type id of ChannelsGetInactiveChannelsRequest.
const ChannelsGetInactiveChannelsRequestTypeID = 0x11e831ee

func (g *ChannelsGetInactiveChannelsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetInactiveChannelsRequest) String() string {
	if g == nil {
		return "ChannelsGetInactiveChannelsRequest(nil)"
	}
	type Alias ChannelsGetInactiveChannelsRequest
	return fmt.Sprintf("ChannelsGetInactiveChannelsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetInactiveChannelsRequest) TypeID() uint32 {
	return ChannelsGetInactiveChannelsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetInactiveChannelsRequest) TypeName() string {
	return "channels.getInactiveChannels"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetInactiveChannelsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getInactiveChannels",
		ID:   ChannelsGetInactiveChannelsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetInactiveChannelsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getInactiveChannels#11e831ee as nil")
	}
	b.PutID(ChannelsGetInactiveChannelsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetInactiveChannelsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getInactiveChannels#11e831ee as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetInactiveChannelsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getInactiveChannels#11e831ee to nil")
	}
	if err := b.ConsumeID(ChannelsGetInactiveChannelsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getInactiveChannels#11e831ee: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetInactiveChannelsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getInactiveChannels#11e831ee to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetInactiveChannelsRequest.
var (
	_ bin.Encoder     = &ChannelsGetInactiveChannelsRequest{}
	_ bin.Decoder     = &ChannelsGetInactiveChannelsRequest{}
	_ bin.BareEncoder = &ChannelsGetInactiveChannelsRequest{}
	_ bin.BareDecoder = &ChannelsGetInactiveChannelsRequest{}
)

// ChannelsGetInactiveChannels invokes method channels.getInactiveChannels#11e831ee returning error if any.
// Get inactive channels and supergroups
//
// See https://core.telegram.org/method/channels.getInactiveChannels for reference.
func (c *Client) ChannelsGetInactiveChannels(ctx context.Context) (*MessagesInactiveChats, error) {
	var result MessagesInactiveChats

	request := &ChannelsGetInactiveChannelsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
