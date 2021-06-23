package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "kir",
		Price: 10000000.0,
		SKU:   "a-b-c",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
