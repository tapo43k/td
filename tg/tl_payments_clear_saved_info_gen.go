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

// PaymentsClearSavedInfoRequest represents TL type `payments.clearSavedInfo#d83d70c1`.
// Clear saved payment information
//
// See https://core.telegram.org/method/payments.clearSavedInfo for reference.
type PaymentsClearSavedInfoRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Remove saved payment credentials
	Credentials bool
	// Clear the last order settings saved by the user
	Info bool
}

// PaymentsClearSavedInfoRequestTypeID is TL type id of PaymentsClearSavedInfoRequest.
const PaymentsClearSavedInfoRequestTypeID = 0xd83d70c1

func (c *PaymentsClearSavedInfoRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Credentials == false) {
		return false
	}
	if !(c.Info == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *PaymentsClearSavedInfoRequest) String() string {
	if c == nil {
		return "PaymentsClearSavedInfoRequest(nil)"
	}
	type Alias PaymentsClearSavedInfoRequest
	return fmt.Sprintf("PaymentsClearSavedInfoRequest%+v", Alias(*c))
}

// FillFrom fills PaymentsClearSavedInfoRequest from given interface.
func (c *PaymentsClearSavedInfoRequest) FillFrom(from interface {
	GetCredentials() (value bool)
	GetInfo() (value bool)
}) {
	c.Credentials = from.GetCredentials()
	c.Info = from.GetInfo()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsClearSavedInfoRequest) TypeID() uint32 {
	return PaymentsClearSavedInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsClearSavedInfoRequest) TypeName() string {
	return "payments.clearSavedInfo"
}

// TypeInfo returns info about TL type.
func (c *PaymentsClearSavedInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.clearSavedInfo",
		ID:   PaymentsClearSavedInfoRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Credentials",
			SchemaName: "credentials",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Info",
			SchemaName: "info",
			Null:       !c.Flags.Has(1),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *PaymentsClearSavedInfoRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.clearSavedInfo#d83d70c1 as nil")
	}
	b.PutID(PaymentsClearSavedInfoRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *PaymentsClearSavedInfoRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode payments.clearSavedInfo#d83d70c1 as nil")
	}
	if !(c.Credentials == false) {
		c.Flags.Set(0)
	}
	if !(c.Info == false) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.clearSavedInfo#d83d70c1: field flags: %w", err)
	}
	return nil
}

// SetCredentials sets value of Credentials conditional field.
func (c *PaymentsClearSavedInfoRequest) SetCredentials(value bool) {
	if value {
		c.Flags.Set(0)
		c.Credentials = true
	} else {
		c.Flags.Unset(0)
		c.Credentials = false
	}
}

// GetCredentials returns value of Credentials conditional field.
func (c *PaymentsClearSavedInfoRequest) GetCredentials() (value bool) {
	return c.Flags.Has(0)
}

// SetInfo sets value of Info conditional field.
func (c *PaymentsClearSavedInfoRequest) SetInfo(value bool) {
	if value {
		c.Flags.Set(1)
		c.Info = true
	} else {
		c.Flags.Unset(1)
		c.Info = false
	}
}

// GetInfo returns value of Info conditional field.
func (c *PaymentsClearSavedInfoRequest) GetInfo() (value bool) {
	return c.Flags.Has(1)
}

// Decode implements bin.Decoder.
func (c *PaymentsClearSavedInfoRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.clearSavedInfo#d83d70c1 to nil")
	}
	if err := b.ConsumeID(PaymentsClearSavedInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.clearSavedInfo#d83d70c1: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *PaymentsClearSavedInfoRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode payments.clearSavedInfo#d83d70c1 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.clearSavedInfo#d83d70c1: field flags: %w", err)
		}
	}
	c.Credentials = c.Flags.Has(0)
	c.Info = c.Flags.Has(1)
	return nil
}

// Ensuring interfaces in compile-time for PaymentsClearSavedInfoRequest.
var (
	_ bin.Encoder     = &PaymentsClearSavedInfoRequest{}
	_ bin.Decoder     = &PaymentsClearSavedInfoRequest{}
	_ bin.BareEncoder = &PaymentsClearSavedInfoRequest{}
	_ bin.BareDecoder = &PaymentsClearSavedInfoRequest{}
)

// PaymentsClearSavedInfo invokes method payments.clearSavedInfo#d83d70c1 returning error if any.
// Clear saved payment information
//
// See https://core.telegram.org/method/payments.clearSavedInfo for reference.
func (c *Client) PaymentsClearSavedInfo(ctx context.Context, request *PaymentsClearSavedInfoRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
