package main

import (
  "testing"
)

// Test for normalizePort.
func TestNormalizePort(t *testing.T) {
  dummyPort := "1000"
  normalized := normalizePort(dummyPort)

  if normalized != ":1000" {
    t.Error("Expected normalized port to be ':1000', got", normalized)
  }
}
