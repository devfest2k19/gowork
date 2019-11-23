package models

import "testing"

func TestGetFullNumber(t *testing.T) {

	p := Person{}
	p.ID = 1
	p.Firstname = "test firstname"
	p.Lastname = "test lastname"
	p.City = "test city"
	p.AreaCode = "045"
	p.Phone = "8475847"

	result := p.GetFullNumber()

	if result != "045-8475847" {
		t.Error("Test Failed GetFullNumber()")
	}
}

func TestGetFullNumberTableDriven(t *testing.T) {
	type TableEntry struct {
		Input  Person
		Output string
	}

	p := Person{}
	p.ID = 1
	p.Firstname = "test firstname"
	p.Lastname = "test lastname"
	p.City = "test city"
	p.AreaCode = "045"
	p.Phone = "8475847"

	p1 := p
	p1.AreaCode = "034"
	p1.Phone = "4751485"

	testData := []TableEntry{
		TableEntry{
			Input:  p,
			Output: "045-8475847",
		},
		TableEntry{
			Input:  p1,
			Output: "034-4751485",
		},
	}

	for _, te := range testData {
		if te.Input.GetFullNumber() != te.Output {
			t.Error("Error GetFullNumber()")
		}
	}
}

func BenchmarkGetFullNumber(b *testing.B) {

	p := Person{}
	p.ID = 1
	p.Firstname = "test firstname"
	p.Lastname = "test lastname"
	p.City = "test city"
	p.AreaCode = "045"
	p.Phone = "8475847"

	for i := 0; i <= b.N; i++ {
		p.GetFullNumber()
	}
}
