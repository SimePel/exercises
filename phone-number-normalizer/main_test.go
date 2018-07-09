package main

import "testing"

func TestNormalize(t *testing.T) {
	phones := [8]string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	expected := [8]string{
		"1234567890",
		"1234567891",
		"1234567892",
		"1234567893",
		"1234567894",
		"1234567890",
		"1234567892",
		"1234567892",
	}

	for i := range phones {
		if expected[i] != Normalize(phones[i]) {
			t.Errorf("Give: %s, Got: %s", phones[i], expected[i])
		}
	}
}
