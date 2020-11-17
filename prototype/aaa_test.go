package prototype

import (
	"testing"
)

func TestClone(t *testing.T) {
	shirtCache := GetShirtsCloner()
	if shirtCache == nil {
		t.Fatalf("?")
	}

	item1, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatalf("?")
	}

	if item1 == whitePrototype {
		t.Fatalf("?")
	}

	shirt1, ok := item1.(*Shirt)
	if !ok {
		t.Fatalf("?")
	}
	shirt1.SKU = "abbcc"
	item2, err := shirtCache.GetClone(White)
	if err != nil {
		t.Fatalf("?")
	}
	shirt2, ok := item2.(*Shirt)
	if !ok {
		t.Fatalf("?")
	}
	if shirt1.SKU == shirt2.SKU {
		t.Fatalf("?")
	}
	if shirt1 == shirt2 {
		t.Fatalf("?")
	}

	t.Logf("LOG: %s", shirt1.GetInfo())
	t.Logf("LOG: %s", shirt2.GetInfo())

	t.Logf("LOG: %p != %p \n\n", &shirt1, &shirt2)
}
