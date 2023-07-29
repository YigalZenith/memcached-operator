//go:build tools
// +build tools

package tools

import (
	// Needed for the storage version too.
	_ "knative.dev/pkg/apiextensions/storageversion/cmd/migrate"
)
