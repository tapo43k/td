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

// MessagesGetChatInviteImportersRequest represents TL type `messages.getChatInviteImporters#26fb7289`.
//
// See https://core.telegram.org/method/messages.getChatInviteImporters for reference.
type MessagesGetChatInviteImportersRequest struct {
	// Peer field of MessagesGetChatInviteImportersRequest.
	Peer InputPeerClass
	// Link field of MessagesGetChatInviteImportersRequest.
	Link string
	// OffsetDate field of MessagesGetChatInviteImportersRequest.
	OffsetDate int
	// OffsetUser field of MessagesGetChatInviteImportersRequest.
	OffsetUser InputUserClass
	// Limit field of MessagesGetChatInviteImportersRequest.
	Limit int
}

// MessagesGetChatInviteImportersRequestTypeID is TL type id of MessagesGetChatInviteImportersRequest.
const MessagesGetChatInviteImportersRequestTypeID = 0x26fb7289

func (g *MessagesGetChatInviteImportersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.Link == "") {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.OffsetUser == nil) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetChatInviteImportersRequest) String() string {
	if g == nil {
		return "MessagesGetChatInviteImportersRequest(nil)"
	}
	type Alias MessagesGetChatInviteImportersRequest
	return fmt.Sprintf("MessagesGetChatInviteImportersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetChatInviteImportersRequest from given interface.
func (g *MessagesGetChatInviteImportersRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetLink() (value string)
	GetOffsetDate() (value int)
	GetOffsetUser() (value InputUserClass)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.Link = from.GetLink()
	g.OffsetDate = from.GetOffsetDate()
	g.OffsetUser = from.GetOffsetUser()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetChatInviteImportersRequest) TypeID() uint32 {
	return MessagesGetChatInviteImportersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetChatInviteImportersRequest) TypeName() string {
	return "messages.getChatInviteImporters"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetChatInviteImportersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getChatInviteImporters",
		ID:   MessagesGetChatInviteImportersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "OffsetUser",
			SchemaName: "offset_user",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetChatInviteImportersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getChatInviteImporters#26fb7289 as nil")
	}
	b.PutID(MessagesGetChatInviteImportersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetChatInviteImportersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getChatInviteImporters#26fb7289 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getChatInviteImporters#26fb7289: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getChatInviteImporters#26fb7289: field peer: %w", err)
	}
	b.PutString(g.Link)
	b.PutInt(g.OffsetDate)
	if g.OffsetUser == nil {
		return fmt.Errorf("unable to encode messages.getChatInviteImporters#26fb7289: field offset_user is nil")
	}
	if err := g.OffsetUser.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getChatInviteImporters#26fb7289: field offset_user: %w", err)
	}
	b.PutInt(g.Limit)
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetChatInviteImportersRequest) GetPeer() (value InputPeerClass) {
	return g.Peer
}

// GetLink returns value of Link field.
func (g *MessagesGetChatInviteImportersRequest) GetLink() (value string) {
	return g.Link
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetChatInviteImportersRequest) GetOffsetDate() (value int) {
	return g.OffsetDate
}

// GetOffsetUser returns value of OffsetUser field.
func (g *MessagesGetChatInviteImportersRequest) GetOffsetUser() (value InputUserClass) {
	return g.OffsetUser
}

// GetLimit returns value of Limit field.
func (g *MessagesGetChatInviteImportersRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *MessagesGetChatInviteImportersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getChatInviteImporters#26fb7289 to nil")
	}
	if err := b.ConsumeID(MessagesGetChatInviteImportersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getChatInviteImporters#26fb7289: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetChatInviteImportersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getChatInviteImporters#26fb7289 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChatInviteImporters#26fb7289: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChatInviteImporters#26fb7289: field link: %w", err)
		}
		g.Link = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChatInviteImporters#26fb7289: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChatInviteImporters#26fb7289: field offset_user: %w", err)
		}
		g.OffsetUser = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getChatInviteImporters#26fb7289: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetChatInviteImportersRequest.
var (
	_ bin.Encoder     = &MessagesGetChatInviteImportersRequest{}
	_ bin.Decoder     = &MessagesGetChatInviteImportersRequest{}
	_ bin.BareEncoder = &MessagesGetChatInviteImportersRequest{}
	_ bin.BareDecoder = &MessagesGetChatInviteImportersRequest{}
)

// MessagesGetChatInviteImporters invokes method messages.getChatInviteImporters#26fb7289 returning error if any.
//
// See https://core.telegram.org/method/messages.getChatInviteImporters for reference.
func (c *Client) MessagesGetChatInviteImporters(ctx context.Context, request *MessagesGetChatInviteImportersRequest) (*MessagesChatInviteImporters, error) {
	var result MessagesChatInviteImporters

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
