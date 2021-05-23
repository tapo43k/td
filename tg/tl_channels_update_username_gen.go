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

// ChannelsUpdateUsernameRequest represents TL type `channels.updateUsername#3514b3de`.
// Change the username of a supergroup/channel
//
// See https://core.telegram.org/method/channels.updateUsername for reference.
type ChannelsUpdateUsernameRequest struct {
	// Channel
	Channel InputChannelClass
	// New username
	Username string
}

// ChannelsUpdateUsernameRequestTypeID is TL type id of ChannelsUpdateUsernameRequest.
const ChannelsUpdateUsernameRequestTypeID = 0x3514b3de

func (u *ChannelsUpdateUsernameRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Channel == nil) {
		return false
	}
	if !(u.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *ChannelsUpdateUsernameRequest) String() string {
	if u == nil {
		return "ChannelsUpdateUsernameRequest(nil)"
	}
	type Alias ChannelsUpdateUsernameRequest
	return fmt.Sprintf("ChannelsUpdateUsernameRequest%+v", Alias(*u))
}

// FillFrom fills ChannelsUpdateUsernameRequest from given interface.
func (u *ChannelsUpdateUsernameRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetUsername() (value string)
}) {
	u.Channel = from.GetChannel()
	u.Username = from.GetUsername()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsUpdateUsernameRequest) TypeID() uint32 {
	return ChannelsUpdateUsernameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsUpdateUsernameRequest) TypeName() string {
	return "channels.updateUsername"
}

// TypeInfo returns info about TL type.
func (u *ChannelsUpdateUsernameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.updateUsername",
		ID:   ChannelsUpdateUsernameRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *ChannelsUpdateUsernameRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateUsername#3514b3de as nil")
	}
	b.PutID(ChannelsUpdateUsernameRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *ChannelsUpdateUsernameRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateUsername#3514b3de as nil")
	}
	if u.Channel == nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel: %w", err)
	}
	b.PutString(u.Username)
	return nil
}

// GetChannel returns value of Channel field.
func (u *ChannelsUpdateUsernameRequest) GetChannel() (value InputChannelClass) {
	return u.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (u *ChannelsUpdateUsernameRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return u.Channel.AsNotEmpty()
}

// GetUsername returns value of Username field.
func (u *ChannelsUpdateUsernameRequest) GetUsername() (value string) {
	return u.Username
}

// Decode implements bin.Decoder.
func (u *ChannelsUpdateUsernameRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateUsername#3514b3de to nil")
	}
	if err := b.ConsumeID(ChannelsUpdateUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *ChannelsUpdateUsernameRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateUsername#3514b3de to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field channel: %w", err)
		}
		u.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsUpdateUsernameRequest.
var (
	_ bin.Encoder     = &ChannelsUpdateUsernameRequest{}
	_ bin.Decoder     = &ChannelsUpdateUsernameRequest{}
	_ bin.BareEncoder = &ChannelsUpdateUsernameRequest{}
	_ bin.BareDecoder = &ChannelsUpdateUsernameRequest{}
)

// ChannelsUpdateUsername invokes method channels.updateUsername#3514b3de returning error if any.
// Change the username of a supergroup/channel
//
// Possible errors:
//  400 CHANNELS_ADMIN_PUBLIC_TOO_MUCH: You're admin of too many public channels, make some channels private to change the username of this channel
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 USERNAME_INVALID: The provided username is not valid
//  400 USERNAME_NOT_MODIFIED: The username was not modified
//  400 USERNAME_OCCUPIED: The provided username is already occupied
//
// See https://core.telegram.org/method/channels.updateUsername for reference.
func (c *Client) ChannelsUpdateUsername(ctx context.Context, request *ChannelsUpdateUsernameRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
