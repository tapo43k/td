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

// ChannelsGetGroupsForDiscussionRequest represents TL type `channels.getGroupsForDiscussion#f5dad378`.
// Get all groups that can be used as discussion groups¹.
// Returned legacy group chats¹ must be first upgraded to supergroups² before they can
// be set as a discussion group.
// To set a returned supergroup as a discussion group, access to its old messages must be
// enabled using channels.togglePreHistoryHidden³, first.
//
// Links:
//  1) https://core.telegram.org/api/discussion
//  2) https://core.telegram.org/api/channel
//  3) https://core.telegram.org/api/channel
//  4) https://core.telegram.org/method/channels.togglePreHistoryHidden
//
// See https://core.telegram.org/method/channels.getGroupsForDiscussion for reference.
type ChannelsGetGroupsForDiscussionRequest struct {
}

// ChannelsGetGroupsForDiscussionRequestTypeID is TL type id of ChannelsGetGroupsForDiscussionRequest.
const ChannelsGetGroupsForDiscussionRequestTypeID = 0xf5dad378

func (g *ChannelsGetGroupsForDiscussionRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetGroupsForDiscussionRequest) String() string {
	if g == nil {
		return "ChannelsGetGroupsForDiscussionRequest(nil)"
	}
	type Alias ChannelsGetGroupsForDiscussionRequest
	return fmt.Sprintf("ChannelsGetGroupsForDiscussionRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetGroupsForDiscussionRequest) TypeID() uint32 {
	return ChannelsGetGroupsForDiscussionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetGroupsForDiscussionRequest) TypeName() string {
	return "channels.getGroupsForDiscussion"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetGroupsForDiscussionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getGroupsForDiscussion",
		ID:   ChannelsGetGroupsForDiscussionRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetGroupsForDiscussionRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getGroupsForDiscussion#f5dad378 as nil")
	}
	b.PutID(ChannelsGetGroupsForDiscussionRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetGroupsForDiscussionRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getGroupsForDiscussion#f5dad378 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetGroupsForDiscussionRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getGroupsForDiscussion#f5dad378 to nil")
	}
	if err := b.ConsumeID(ChannelsGetGroupsForDiscussionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getGroupsForDiscussion#f5dad378: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetGroupsForDiscussionRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getGroupsForDiscussion#f5dad378 to nil")
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsGetGroupsForDiscussionRequest.
var (
	_ bin.Encoder     = &ChannelsGetGroupsForDiscussionRequest{}
	_ bin.Decoder     = &ChannelsGetGroupsForDiscussionRequest{}
	_ bin.BareEncoder = &ChannelsGetGroupsForDiscussionRequest{}
	_ bin.BareDecoder = &ChannelsGetGroupsForDiscussionRequest{}
)

// ChannelsGetGroupsForDiscussion invokes method channels.getGroupsForDiscussion#f5dad378 returning error if any.
// Get all groups that can be used as discussion groups¹.
// Returned legacy group chats¹ must be first upgraded to supergroups² before they can
// be set as a discussion group.
// To set a returned supergroup as a discussion group, access to its old messages must be
// enabled using channels.togglePreHistoryHidden³, first.
//
// Links:
//  1) https://core.telegram.org/api/discussion
//  2) https://core.telegram.org/api/channel
//  3) https://core.telegram.org/api/channel
//  4) https://core.telegram.org/method/channels.togglePreHistoryHidden
//
// See https://core.telegram.org/method/channels.getGroupsForDiscussion for reference.
func (c *Client) ChannelsGetGroupsForDiscussion(ctx context.Context) (MessagesChatsClass, error) {
	var result MessagesChatsBox

	request := &ChannelsGetGroupsForDiscussionRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Chats, nil
}
