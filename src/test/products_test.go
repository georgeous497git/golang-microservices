package test

import (
	"GoMicroservices/models"
	"testing"
)

func TestCheckValidation(t *testing.T) {

	p := &models.Product{
		Name:  "Coffee Test",
		Price: 10.00,
		Sku:   "abc-123",
	}

	error := p.Validate()

	if error != nil {
		t.Fatal("Expected error", error)
	}
}
