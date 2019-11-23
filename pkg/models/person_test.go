package models

import "testing"

func TestGetFullNumber(t *testing.T) {
	per := Person{}
	per.FirstName = "fname"
	per.LastName = "lname"
	per.Id = 1
	per.AreaCode = "045"
	per.Phone = "5432345"

	number := per.GetFullNumber()

	if number != "045-5432345" {
		t.Error("Error: GetFullNumber()")
	}
}

func BenchmarkGetFullNumber(b *testing.B) {
	per := Person{}
	per.FirstName = "fname"
	per.LastName = "lname"
	per.Id = 1
	per.AreaCode = "045"
	per.Phone = "5432345"

	for i := 0; i <= b.N; i++ {
		per.GetFullNumber()
	}
}
