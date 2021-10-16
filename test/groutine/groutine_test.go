package groutinet

import (
	"testing"

	groutinet "photowey.com/the-grammar-to-go/grammar/groutine"
)

func TestHandleGroutine(t *testing.T) {
	groutinet.HandleGroutine(5)
}

func TestHandleGroutineError(t *testing.T) {
	groutinet.HandleGroutineError(5)
}
