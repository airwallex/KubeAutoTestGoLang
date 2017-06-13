package main

import (
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	funResponse := string("BANANA")

	if !strings.Contains(funResponse, Message()) {
		t.Errorf("Expecting %v but got %v", funResponse, Message())
	}
}
