// SPDX-License-Identifier: MIT

package admindb

import (
	"context"

	"go.mindeco.de/http/auth"
	refs "go.mindeco.de/ssb-refs"
)

// AuthFallbackService might be helpful for scenarios where one lost access to his ssb device or key
type AuthFallbackService interface {
	auth.Auther

	Create(ctx context.Context, user string, password []byte) error
	GetByID(ctx context.Context, uid int64) (*User, error)
}

// AuthWithSSBService defines functions needed for the challange/response system of sign-in with ssb
type AuthWithSSBService interface{}

// AllowListService changes the lists of people that are allowed to get into the room
type AllowListService interface {
	// Add adds the feed to the list.
	Add(context.Context, refs.FeedRef) error

	// HasFeed returns true if a feed is on the list.
	HasFeed(context.Context, refs.FeedRef) bool

	// HasFeed returns true if a feed is on the list.
	HasID(context.Context, int64) bool

	// GetByID returns the list entry for that ID or an error
	GetByID(context.Context, int64) (ListEntry, error)

	// List returns a list of all the feeds.
	List(context.Context) (ListEntries, error)

	// RemoveFeed removes the feed from the list.
	RemoveFeed(context.Context, refs.FeedRef) error

	// RemoveID removes the feed for the ID from the list.
	RemoveID(context.Context, int64) error
}

// AliasService manages alias handle registration and lookup
type AliasService interface{}

// for tests we use generated mocks from these interfaces created with https://github.com/maxbrunsfeld/counterfeiter

//go:generate counterfeiter -o mockdb/auth.go . AuthWithSSBService

//go:generate counterfeiter -o mockdb/auth_fallback.go . AuthFallbackService

//go:generate counterfeiter -o mockdb/allow.go . AllowListService

//go:generate counterfeiter -o mockdb/alias.go . AliasService