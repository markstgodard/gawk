package main

import (
	"testing"
)

func TestNewServer(t *testing.T) {

	args := []string{"logs"}
	_, err := newServer(args)
	if err != nil {
		t.Errorf("Error creating newServer", err)
	}

}
