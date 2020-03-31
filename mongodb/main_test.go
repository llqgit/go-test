package main

import (
	"testing"
)

func TestInlineStructInsert(t *testing.T) {
	InlineStructInsert()
}

func TestInlineStructFetch(t *testing.T) {
	list := InlineStructFetch()
	t.Log(list)
}
