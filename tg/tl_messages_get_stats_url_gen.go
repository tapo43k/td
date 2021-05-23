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

// MessagesGetStatsURLRequest represents TL type `messages.getStatsURL#812c2ae6`.
// Returns URL with the chat statistics. Currently this method can be used only for
// channels
//
// See https://core.telegram.org/method/messages.getStatsURL for reference.
type MessagesGetStatsURLRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Pass true if a URL with the dark theme must be returned
	Dark bool
	// Chat identifier
	Peer InputPeerClass
	// Parameters from tg://statsrefresh?params=****** link
	Params string
}

// MessagesGetStatsURLRequestTypeID is TL type id of MessagesGetStatsURLRequest.
const MessagesGetStatsURLRequestTypeID = 0x812c2ae6

func (g *MessagesGetStatsURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Dark == false) {
		return false
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Params == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetStatsURLRequest) String() string {
	if g == nil {
		return "MessagesGetStatsURLRequest(nil)"
	}
	type Alias MessagesGetStatsURLRequest
	return fmt.Sprintf("MessagesGetStatsURLRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetStatsURLRequest from given interface.
func (g *MessagesGetStatsURLRequest) FillFrom(from interface {
	GetDark() (value bool)
	GetPeer() (value InputPeerClass)
	GetParams() (value string)
}) {
	g.Dark = from.GetDark()
	g.Peer = from.GetPeer()
	g.Params = from.GetParams()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetStatsURLRequest) TypeID() uint32 {
	return MessagesGetStatsURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetStatsURLRequest) TypeName() string {
	return "messages.getStatsURL"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetStatsURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getStatsURL",
		ID:   MessagesGetStatsURLRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Params",
			SchemaName: "params",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetStatsURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStatsURL#812c2ae6 as nil")
	}
	b.PutID(MessagesGetStatsURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetStatsURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getStatsURL#812c2ae6 as nil")
	}
	if !(g.Dark == false) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStatsURL#812c2ae6: field flags: %w", err)
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getStatsURL#812c2ae6: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getStatsURL#812c2ae6: field peer: %w", err)
	}
	b.PutString(g.Params)
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *MessagesGetStatsURLRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
		g.Dark = true
	} else {
		g.Flags.Unset(0)
		g.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (g *MessagesGetStatsURLRequest) GetDark() (value bool) {
	return g.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (g *MessagesGetStatsURLRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetParams returns value of Params field.
func (g *MessagesGetStatsURLRequest) GetParams() (value string) {
	return g.Params
}

// Decode implements bin.Decoder.
func (g *MessagesGetStatsURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStatsURL#812c2ae6 to nil")
	}
	if err := b.ConsumeID(MessagesGetStatsURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetStatsURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getStatsURL#812c2ae6 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getStatsURL#812c2ae6: field params: %w", err)
		}
		g.Params = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetStatsURLRequest.
var (
	_ bin.Encoder     = &MessagesGetStatsURLRequest{}
	_ bin.Decoder     = &MessagesGetStatsURLRequest{}
	_ bin.BareEncoder = &MessagesGetStatsURLRequest{}
	_ bin.BareDecoder = &MessagesGetStatsURLRequest{}
)

// MessagesGetStatsURL invokes method messages.getStatsURL#812c2ae6 returning error if any.
// Returns URL with the chat statistics. Currently this method can be used only for
// channels
//
// Possible errors:
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getStatsURL for reference.
func (c *Client) MessagesGetStatsURL(ctx context.Context, request *MessagesGetStatsURLRequest) (*StatsURL, error) {
	var result StatsURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
