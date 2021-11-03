package test

import (
	"data"
	"testing"
)

func TestCheckValidation(t *testing.T) {

	p := &data.Product{
		Name:  "Coffee Test",
		Price: 10.00,
		Sku:   "abc-123",
	}

	error := p.Validate()

	if error != nil {
		t.Fatal("Expected error", error)
	}
}
