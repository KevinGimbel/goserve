// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Run tests
func Test() error {
	return sh.Run("go", "test")
}

// Create a full release with goreleaser (Pushes to GitHub)
func Release() error {
	return sh.Run("goreleaser", "--rm-dist")
}

// Create a pre-release with goreleaser
func PreRelease() error {
	return sh.Run("goreleaser", "--rm-dist", "--snapshot")
}
