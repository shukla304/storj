// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

package rewards

import (
	"context"
	"time"

	"github.com/zeebo/errs"

	"storj.io/storj/private/currency"
)

var (
	// ErrReachedMaxCapacity is the error class used when an offer has reached its redemption capacity.
	ErrReachedMaxCapacity = errs.Class("offer redemption has reached its capacity")
	// ErrOfferNotExist is the error class used when no current offer is set.
	ErrOfferNotExist = errs.Class("no current offer")
)

// DB holds information about offer
//
// architecture: Database
type DB interface {
	ListAll(ctx context.Context) (Offers, error)
	GetActiveOffersByType(ctx context.Context, offerType OfferType) (Offers, error)
	Create(ctx context.Context, offer *NewOffer) (*Offer, error)
	Finish(ctx context.Context, offerID int) error
}

// NewOffer holds information that's needed for creating a new offer
type NewOffer struct {
	Name        string
	Description string

	AwardCredit   currency.USD
	InviteeCredit currency.USD

	RedeemableCap int

	AwardCreditDurationDays   int
	InviteeCreditDurationDays int

	ExpiresAt time.Time

	Status OfferStatus
	Type   OfferType
}

// UpdateOffer holds fields needed for update an offer
type UpdateOffer struct {
	ID        int
	Status    OfferStatus
	ExpiresAt time.Time
}

// RedeemOffer holds field needed for redeem an offer
type RedeemOffer struct {
	RedeemableCap int
	Status        OfferStatus
	Type          OfferType
}

// Offers contains a slice of offers.
type Offers []Offer

// OfferType indicates the type of an offer
type OfferType int

const (
	// Invalid is a default value for offers that don't have correct type associated with it
	Invalid = OfferType(0)
	// FreeCredit is a type of offers used for Free Credit Program
	FreeCredit = OfferType(1)
	// Referral is a type of offers used for Referral Program
	Referral = OfferType(2)
	// Partner is an OfferType used be the Open Source Partner Program
	Partner = OfferType(3)
)

// OfferStatus represents the different stage an offer can have in its life-cycle.
type OfferStatus int

const (
	// Done is the status of an offer that is no longer in use.
	Done = OfferStatus(iota)
	// Default is the status of an offer when there is no active offer.
	Default
	// Active is the status of an offer that is currently in use.
	Active
)

// Offer contains info needed for giving users free credits through different offer programs
type Offer struct {
	ID          int
	Name        string
	Description string

	AwardCredit   currency.USD
	InviteeCredit currency.USD

	AwardCreditDurationDays   int
	InviteeCreditDurationDays int

	RedeemableCap int

	ExpiresAt time.Time
	CreatedAt time.Time

	Status OfferStatus
	Type   OfferType
}

// IsEmpty evaluates whether or not an on offer is empty.
func (offer Offer) IsEmpty() bool {
	return offer.Name == ""
}

// IsZero returns whether it's equivalent to empty struct.
func (offer Offer) IsZero() bool {
	return offer == Offer{}
}

// IsDefault checks if a offer's status is default
func (status OfferStatus) IsDefault() bool {
	return status == Default
}
