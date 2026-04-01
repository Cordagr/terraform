// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: BUSL-1.1

package getproviders

import (
	"context"
	"time"

	"github.com/hashicorp/terraform/internal/addrs"
)

// A Source can query a particular source for information about providers
// that are available to install.
type Source interface {
	AvailableVersions(ctx context.Context, provider addrs.Provider) (VersionList, Warnings, error)
	PackageMeta(ctx context.Context, provider addrs.Provider, version Version, target Platform) (PackageMeta, error)
	ForDisplay(provider addrs.Provider) string
}

// VersionTimestampSource is an optional extension for Source implementations
// that can report publish timestamps for provider versions.
type VersionTimestampSource interface {
	VersionTimestamp(ctx context.Context, provider addrs.Provider, version Version) (*time.Time, error)
}
