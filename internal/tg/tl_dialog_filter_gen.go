// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// DialogFilter represents TL type `dialogFilter#7438f7e8`.
type DialogFilter struct {
	// Flags field of DialogFilter.
	Flags bin.Fields
	// Contacts field of DialogFilter.
	Contacts bool
	// NonContacts field of DialogFilter.
	NonContacts bool
	// Groups field of DialogFilter.
	Groups bool
	// Broadcasts field of DialogFilter.
	Broadcasts bool
	// Bots field of DialogFilter.
	Bots bool
	// ExcludeMuted field of DialogFilter.
	ExcludeMuted bool
	// ExcludeRead field of DialogFilter.
	ExcludeRead bool
	// ExcludeArchived field of DialogFilter.
	ExcludeArchived bool
	// ID field of DialogFilter.
	ID int
	// Title field of DialogFilter.
	Title string
	// Emoticon field of DialogFilter.
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string
	// PinnedPeers field of DialogFilter.
	PinnedPeers []InputPeerClass
	// IncludePeers field of DialogFilter.
	IncludePeers []InputPeerClass
	// ExcludePeers field of DialogFilter.
	ExcludePeers []InputPeerClass
}

// DialogFilterTypeID is TL type id of DialogFilter.
const DialogFilterTypeID = 0x7438f7e8

// Encode implements bin.Encoder.
func (d *DialogFilter) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilter#7438f7e8 as nil")
	}
	b.PutID(DialogFilterTypeID)
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.Title)
	if d.Flags.Has(25) {
		b.PutString(d.Emoticon)
	}
	b.PutVectorHeader(len(d.PinnedPeers))
	for idx, v := range d.PinnedPeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field pinned_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field pinned_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.IncludePeers))
	for idx, v := range d.IncludePeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field include_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field include_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.ExcludePeers))
	for idx, v := range d.ExcludePeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field exclude_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilter#7438f7e8: field exclude_peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetContacts sets value of Contacts conditional field.
func (d *DialogFilter) SetContacts(value bool) {
	if value {
		d.Flags.Set(0)
	} else {
		d.Flags.Unset(0)
	}
}

// SetNonContacts sets value of NonContacts conditional field.
func (d *DialogFilter) SetNonContacts(value bool) {
	if value {
		d.Flags.Set(1)
	} else {
		d.Flags.Unset(1)
	}
}

// SetGroups sets value of Groups conditional field.
func (d *DialogFilter) SetGroups(value bool) {
	if value {
		d.Flags.Set(2)
	} else {
		d.Flags.Unset(2)
	}
}

// SetBroadcasts sets value of Broadcasts conditional field.
func (d *DialogFilter) SetBroadcasts(value bool) {
	if value {
		d.Flags.Set(3)
	} else {
		d.Flags.Unset(3)
	}
}

// SetBots sets value of Bots conditional field.
func (d *DialogFilter) SetBots(value bool) {
	if value {
		d.Flags.Set(4)
	} else {
		d.Flags.Unset(4)
	}
}

// SetExcludeMuted sets value of ExcludeMuted conditional field.
func (d *DialogFilter) SetExcludeMuted(value bool) {
	if value {
		d.Flags.Set(11)
	} else {
		d.Flags.Unset(11)
	}
}

// SetExcludeRead sets value of ExcludeRead conditional field.
func (d *DialogFilter) SetExcludeRead(value bool) {
	if value {
		d.Flags.Set(12)
	} else {
		d.Flags.Unset(12)
	}
}

// SetExcludeArchived sets value of ExcludeArchived conditional field.
func (d *DialogFilter) SetExcludeArchived(value bool) {
	if value {
		d.Flags.Set(13)
	} else {
		d.Flags.Unset(13)
	}
}

// SetEmoticon sets value of Emoticon conditional field.
func (d *DialogFilter) SetEmoticon(value string) {
	d.Flags.Set(25)
	d.Emoticon = value
}

// GetEmoticon returns value of Emoticon conditional field and
// boolean which is true if field was set.
func (d *DialogFilter) GetEmoticon() (value string, ok bool) {
	if !d.Flags.Has(25) {
		return value, false
	}
	return d.Emoticon, true
}

// Decode implements bin.Decoder.
func (d *DialogFilter) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilter#7438f7e8 to nil")
	}
	if err := b.ConsumeID(DialogFilterTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFilter#7438f7e8: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field flags: %w", err)
		}
	}
	d.Contacts = d.Flags.Has(0)
	d.NonContacts = d.Flags.Has(1)
	d.Groups = d.Flags.Has(2)
	d.Broadcasts = d.Flags.Has(3)
	d.Bots = d.Flags.Has(4)
	d.ExcludeMuted = d.Flags.Has(11)
	d.ExcludeRead = d.Flags.Has(12)
	d.ExcludeArchived = d.Flags.Has(13)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field title: %w", err)
		}
		d.Title = value
	}
	if d.Flags.Has(25) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field emoticon: %w", err)
		}
		d.Emoticon = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field pinned_peers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field pinned_peers: %w", err)
			}
			d.PinnedPeers = append(d.PinnedPeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field include_peers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field include_peers: %w", err)
			}
			d.IncludePeers = append(d.IncludePeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field exclude_peers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilter#7438f7e8: field exclude_peers: %w", err)
			}
			d.ExcludePeers = append(d.ExcludePeers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for DialogFilter.
var (
	_ bin.Encoder = &DialogFilter{}
	_ bin.Decoder = &DialogFilter{}
)