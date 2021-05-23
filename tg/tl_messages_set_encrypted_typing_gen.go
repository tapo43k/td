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

// MessagesSetEncryptedTypingRequest represents TL type `messages.setEncryptedTyping#791451ed`.
// Send typing event by the current user to a secret chat.
//
// See https://core.telegram.org/method/messages.setEncryptedTyping for reference.
type MessagesSetEncryptedTypingRequest struct {
	// Secret chat ID
	Peer InputEncryptedChat
	// Typing.Possible values:(boolTrue)¹, if the user started typing and more than 5
	// seconds have passed since the last request(boolFalse)², if the user stopped typing
	//
	// Links:
	//  1) https://core.telegram.org/constructor/boolTrue
	//  2) https://core.telegram.org/constructor/boolFalse
	Typing bool
}

// MessagesSetEncryptedTypingRequestTypeID is TL type id of MessagesSetEncryptedTypingRequest.
const MessagesSetEncryptedTypingRequestTypeID = 0x791451ed

func (s *MessagesSetEncryptedTypingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Typing == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetEncryptedTypingRequest) String() string {
	if s == nil {
		return "MessagesSetEncryptedTypingRequest(nil)"
	}
	type Alias MessagesSetEncryptedTypingRequest
	return fmt.Sprintf("MessagesSetEncryptedTypingRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetEncryptedTypingRequest from given interface.
func (s *MessagesSetEncryptedTypingRequest) FillFrom(from interface {
	GetPeer() (value InputEncryptedChat)
	GetTyping() (value bool)
}) {
	s.Peer = from.GetPeer()
	s.Typing = from.GetTyping()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetEncryptedTypingRequest) TypeID() uint32 {
	return MessagesSetEncryptedTypingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetEncryptedTypingRequest) TypeName() string {
	return "messages.setEncryptedTyping"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetEncryptedTypingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setEncryptedTyping",
		ID:   MessagesSetEncryptedTypingRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Typing",
			SchemaName: "typing",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetEncryptedTypingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setEncryptedTyping#791451ed as nil")
	}
	b.PutID(MessagesSetEncryptedTypingRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetEncryptedTypingRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setEncryptedTyping#791451ed as nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setEncryptedTyping#791451ed: field peer: %w", err)
	}
	b.PutBool(s.Typing)
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetEncryptedTypingRequest) GetPeer() (value InputEncryptedChat) {
	return s.Peer
}

// GetTyping returns value of Typing field.
func (s *MessagesSetEncryptedTypingRequest) GetTyping() (value bool) {
	return s.Typing
}

// Decode implements bin.Decoder.
func (s *MessagesSetEncryptedTypingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setEncryptedTyping#791451ed to nil")
	}
	if err := b.ConsumeID(MessagesSetEncryptedTypingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetEncryptedTypingRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setEncryptedTyping#791451ed to nil")
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: field peer: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: field typing: %w", err)
		}
		s.Typing = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetEncryptedTypingRequest.
var (
	_ bin.Encoder     = &MessagesSetEncryptedTypingRequest{}
	_ bin.Decoder     = &MessagesSetEncryptedTypingRequest{}
	_ bin.BareEncoder = &MessagesSetEncryptedTypingRequest{}
	_ bin.BareDecoder = &MessagesSetEncryptedTypingRequest{}
)

// MessagesSetEncryptedTyping invokes method messages.setEncryptedTyping#791451ed returning error if any.
// Send typing event by the current user to a secret chat.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//
// See https://core.telegram.org/method/messages.setEncryptedTyping for reference.
func (c *Client) MessagesSetEncryptedTyping(ctx context.Context, request *MessagesSetEncryptedTypingRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
