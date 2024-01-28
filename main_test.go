package main

import (
	"testing"

	"decoder"
	"encoder"
)

var tests = []string{
	"iam",
	"i am",
	"Hi there, how's it going?",
	"amr_halim2008@yahoo.com",
	"you are old",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et 1234!@#$%^&*()_+ dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in 56789-98765 reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
}

func Test3310(t *testing.T) {
	for _, tc := range tests {
		t.Run("Should decode and encode correctly", func(t *testing.T) {
			encoded := encoder.Encode(tc)
			decoded := decoder.Decode(encoded)
			if decoded != tc {
				t.Fatalf("Didn't decode correctly!\nGiven: '%s'.\nEncoded to: '%s'.\nDecoded to: '%s'", tc, encoded, decoded)
			}
		})
	}
}
