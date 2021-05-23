// Code generated by gotdgen, DO NOT EDIT.

package td

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

// DoAuthRequest represents TL type `doAuth#fd2f6687`.
//
// See https://localhost:80/doc/method/doAuth for reference.
type DoAuthRequest struct {
}

// DoAuthRequestTypeID is TL type id of DoAuthRequest.
const DoAuthRequestTypeID = 0xfd2f6687

func (d *DoAuthRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DoAuthRequest) String() string {
	if d == nil {
		return "DoAuthRequest(nil)"
	}
	type Alias DoAuthRequest
	return fmt.Sprintf("DoAuthRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DoAuthRequest) TypeID() uint32 {
	return DoAuthRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DoAuthRequest) TypeName() string {
	return "doAuth"
}

// TypeInfo returns info about TL type.
func (d *DoAuthRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "doAuth",
		ID:   DoAuthRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *DoAuthRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode doAuth#fd2f6687 as nil")
	}
	b.PutID(DoAuthRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DoAuthRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode doAuth#fd2f6687 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DoAuthRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode doAuth#fd2f6687 to nil")
	}
	if err := b.ConsumeID(DoAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode doAuth#fd2f6687: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DoAuthRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode doAuth#fd2f6687 to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for DoAuthRequest.
var (
	_ bin.Encoder     = &DoAuthRequest{}
	_ bin.Decoder     = &DoAuthRequest{}
	_ bin.BareEncoder = &DoAuthRequest{}
	_ bin.BareDecoder = &DoAuthRequest{}
)

// DoAuth invokes method doAuth#fd2f6687 returning error if any.
//
// See https://localhost:80/doc/method/doAuth for reference.
func (c *Client) DoAuth(ctx context.Context) (AuthClass, error) {
	var result AuthBox

	request := &DoAuthRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Auth, nil
}
