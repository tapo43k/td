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

// MessagesGetCommonChatsRequest represents TL type `messages.getCommonChats#d0a48c4`.
// Get chats in common with a user
//
// See https://core.telegram.org/method/messages.getCommonChats for reference.
type MessagesGetCommonChatsRequest struct {
	// User ID
	UserID InputUserClass
	// Maximum ID of chat to return (see pagination¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MaxID int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// MessagesGetCommonChatsRequestTypeID is TL type id of MessagesGetCommonChatsRequest.
const MessagesGetCommonChatsRequestTypeID = 0xd0a48c4

func (g *MessagesGetCommonChatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == nil) {
		return false
	}
	if !(g.MaxID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetCommonChatsRequest) String() string {
	if g == nil {
		return "MessagesGetCommonChatsRequest(nil)"
	}
	type Alias MessagesGetCommonChatsRequest
	return fmt.Sprintf("MessagesGetCommonChatsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetCommonChatsRequest from given interface.
func (g *MessagesGetCommonChatsRequest) FillFrom(from interface {
	GetUserID() (value InputUserClass)
	GetMaxID() (value int)
	GetLimit() (value int)
}) {
	g.UserID = from.GetUserID()
	g.MaxID = from.GetMaxID()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetCommonChatsRequest) TypeID() uint32 {
	return MessagesGetCommonChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetCommonChatsRequest) TypeName() string {
	return "messages.getCommonChats"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetCommonChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getCommonChats",
		ID:   MessagesGetCommonChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetCommonChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getCommonChats#d0a48c4 as nil")
	}
	b.PutID(MessagesGetCommonChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetCommonChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getCommonChats#d0a48c4 as nil")
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode messages.getCommonChats#d0a48c4: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getCommonChats#d0a48c4: field user_id: %w", err)
	}
	b.PutInt(g.MaxID)
	b.PutInt(g.Limit)
	return nil
}

// GetUserID returns value of UserID field.
func (g *MessagesGetCommonChatsRequest) GetUserID() (value InputUserClass) {
	return g.UserID
}

// GetMaxID returns value of MaxID field.
func (g *MessagesGetCommonChatsRequest) GetMaxID() (value int) {
	return g.MaxID
}

// GetLimit returns value of Limit field.
func (g *MessagesGetCommonChatsRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *MessagesGetCommonChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getCommonChats#d0a48c4 to nil")
	}
	if err := b.ConsumeID(MessagesGetCommonChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetCommonChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getCommonChats#d0a48c4 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: field user_id: %w", err)
		}
		g.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: field max_id: %w", err)
		}
		g.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getCommonChats#d0a48c4: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetCommonChatsRequest.
var (
	_ bin.Encoder     = &MessagesGetCommonChatsRequest{}
	_ bin.Decoder     = &MessagesGetCommonChatsRequest{}
	_ bin.BareEncoder = &MessagesGetCommonChatsRequest{}
	_ bin.BareDecoder = &MessagesGetCommonChatsRequest{}
)

// MessagesGetCommonChats invokes method messages.getCommonChats#d0a48c4 returning error if any.
// Get chats in common with a user
//
// Possible errors:
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/messages.getCommonChats for reference.
func (c *Client) MessagesGetCommonChats(ctx context.Context, request *MessagesGetCommonChatsRequest) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
