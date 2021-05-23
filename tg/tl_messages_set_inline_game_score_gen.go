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

// MessagesSetInlineGameScoreRequest represents TL type `messages.setInlineGameScore#15ad9f64`.
// Use this method to set the score of the specified user in a game sent as an inline
// message (bots only).
//
// See https://core.telegram.org/method/messages.setInlineGameScore for reference.
type MessagesSetInlineGameScoreRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag if the game message should be automatically edited to include the
	// current scoreboard
	EditMessage bool
	// Set this flag if the high score is allowed to decrease. This can be useful when fixing
	// mistakes or banning cheaters
	Force bool
	// ID of the inline message
	ID InputBotInlineMessageID
	// User identifier
	UserID InputUserClass
	// New score
	Score int
}

// MessagesSetInlineGameScoreRequestTypeID is TL type id of MessagesSetInlineGameScoreRequest.
const MessagesSetInlineGameScoreRequestTypeID = 0x15ad9f64

func (s *MessagesSetInlineGameScoreRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.EditMessage == false) {
		return false
	}
	if !(s.Force == false) {
		return false
	}
	if !(s.ID.Zero()) {
		return false
	}
	if !(s.UserID == nil) {
		return false
	}
	if !(s.Score == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetInlineGameScoreRequest) String() string {
	if s == nil {
		return "MessagesSetInlineGameScoreRequest(nil)"
	}
	type Alias MessagesSetInlineGameScoreRequest
	return fmt.Sprintf("MessagesSetInlineGameScoreRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetInlineGameScoreRequest from given interface.
func (s *MessagesSetInlineGameScoreRequest) FillFrom(from interface {
	GetEditMessage() (value bool)
	GetForce() (value bool)
	GetID() (value InputBotInlineMessageID)
	GetUserID() (value InputUserClass)
	GetScore() (value int)
}) {
	s.EditMessage = from.GetEditMessage()
	s.Force = from.GetForce()
	s.ID = from.GetID()
	s.UserID = from.GetUserID()
	s.Score = from.GetScore()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetInlineGameScoreRequest) TypeID() uint32 {
	return MessagesSetInlineGameScoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetInlineGameScoreRequest) TypeName() string {
	return "messages.setInlineGameScore"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetInlineGameScoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setInlineGameScore",
		ID:   MessagesSetInlineGameScoreRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EditMessage",
			SchemaName: "edit_message",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Force",
			SchemaName: "force",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSetInlineGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineGameScore#15ad9f64 as nil")
	}
	b.PutID(MessagesSetInlineGameScoreRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetInlineGameScoreRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineGameScore#15ad9f64 as nil")
	}
	if !(s.EditMessage == false) {
		s.Flags.Set(0)
	}
	if !(s.Force == false) {
		s.Flags.Set(1)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field flags: %w", err)
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field id: %w", err)
	}
	if s.UserID == nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field user_id is nil")
	}
	if err := s.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field user_id: %w", err)
	}
	b.PutInt(s.Score)
	return nil
}

// SetEditMessage sets value of EditMessage conditional field.
func (s *MessagesSetInlineGameScoreRequest) SetEditMessage(value bool) {
	if value {
		s.Flags.Set(0)
		s.EditMessage = true
	} else {
		s.Flags.Unset(0)
		s.EditMessage = false
	}
}

// GetEditMessage returns value of EditMessage conditional field.
func (s *MessagesSetInlineGameScoreRequest) GetEditMessage() (value bool) {
	return s.Flags.Has(0)
}

// SetForce sets value of Force conditional field.
func (s *MessagesSetInlineGameScoreRequest) SetForce(value bool) {
	if value {
		s.Flags.Set(1)
		s.Force = true
	} else {
		s.Flags.Unset(1)
		s.Force = false
	}
}

// GetForce returns value of Force conditional field.
func (s *MessagesSetInlineGameScoreRequest) GetForce() (value bool) {
	return s.Flags.Has(1)
}

// GetID returns value of ID field.
func (s *MessagesSetInlineGameScoreRequest) GetID() (value InputBotInlineMessageID) {
	return s.ID
}

// GetUserID returns value of UserID field.
func (s *MessagesSetInlineGameScoreRequest) GetUserID() (value InputUserClass) {
	return s.UserID
}

// GetScore returns value of Score field.
func (s *MessagesSetInlineGameScoreRequest) GetScore() (value int) {
	return s.Score
}

// Decode implements bin.Decoder.
func (s *MessagesSetInlineGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineGameScore#15ad9f64 to nil")
	}
	if err := b.ConsumeID(MessagesSetInlineGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetInlineGameScoreRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineGameScore#15ad9f64 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field flags: %w", err)
		}
	}
	s.EditMessage = s.Flags.Has(0)
	s.Force = s.Flags.Has(1)
	{
		if err := s.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field id: %w", err)
		}
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field score: %w", err)
		}
		s.Score = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetInlineGameScoreRequest.
var (
	_ bin.Encoder     = &MessagesSetInlineGameScoreRequest{}
	_ bin.Decoder     = &MessagesSetInlineGameScoreRequest{}
	_ bin.BareEncoder = &MessagesSetInlineGameScoreRequest{}
	_ bin.BareDecoder = &MessagesSetInlineGameScoreRequest{}
)

// MessagesSetInlineGameScore invokes method messages.setInlineGameScore#15ad9f64 returning error if any.
// Use this method to set the score of the specified user in a game sent as an inline
// message (bots only).
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 USER_BOT_REQUIRED: This method can only be called by a bot
//
// See https://core.telegram.org/method/messages.setInlineGameScore for reference.
// Can be used by bots.
func (c *Client) MessagesSetInlineGameScore(ctx context.Context, request *MessagesSetInlineGameScoreRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
