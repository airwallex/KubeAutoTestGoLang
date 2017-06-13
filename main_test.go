package main

import (
	"strings"
	"testing"
)

func TestMessage(t *testing.T) {
	funResponse := string("Banana")

	if !strings.Contains(funResponse, Message()) {
		t.Errorf("Expecting %v but got %v", funResponse, Message())
	}
}
