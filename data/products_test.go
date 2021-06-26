package data

import "testing"

func TestChecksValidation(t *testing.T) {
	v := NewValidation()
	p := &Product{
		Name:  "kir",
		Price: 10000000.0,
		SKU:   "a-b-c",
	}

	err := v.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
