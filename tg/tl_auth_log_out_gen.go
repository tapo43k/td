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

// AuthLogOutRequest represents TL type `auth.logOut#5717da40`.
// Logs out the user.
//
// See https://core.telegram.org/method/auth.logOut for reference.
type AuthLogOutRequest struct {
}

// AuthLogOutRequestTypeID is TL type id of AuthLogOutRequest.
const AuthLogOutRequestTypeID = 0x5717da40

func (l *AuthLogOutRequest) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *AuthLogOutRequest) String() string {
	if l == nil {
		return "AuthLogOutRequest(nil)"
	}
	type Alias AuthLogOutRequest
	return fmt.Sprintf("AuthLogOutRequest%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthLogOutRequest) TypeID() uint32 {
	return AuthLogOutRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthLogOutRequest) TypeName() string {
	return "auth.logOut"
}

// TypeInfo returns info about TL type.
func (l *AuthLogOutRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.logOut",
		ID:   AuthLogOutRequestTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (l *AuthLogOutRequest) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.logOut#5717da40 as nil")
	}
	b.PutID(AuthLogOutRequestTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *AuthLogOutRequest) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.logOut#5717da40 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLogOutRequest) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.logOut#5717da40 to nil")
	}
	if err := b.ConsumeID(AuthLogOutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.logOut#5717da40: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *AuthLogOutRequest) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.logOut#5717da40 to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for AuthLogOutRequest.
var (
	_ bin.Encoder     = &AuthLogOutRequest{}
	_ bin.Decoder     = &AuthLogOutRequest{}
	_ bin.BareEncoder = &AuthLogOutRequest{}
	_ bin.BareDecoder = &AuthLogOutRequest{}
)

// AuthLogOut invokes method auth.logOut#5717da40 returning error if any.
// Logs out the user.
//
// See https://core.telegram.org/method/auth.logOut for reference.
// Can be used by bots.
func (c *Client) AuthLogOut(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AuthLogOutRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
