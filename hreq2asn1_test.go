package hreq2asn1_test

import (
	"testing"

	qa "github.com/takanoriyanagitani/go-http-req2asn1"
)

func TestAsn1Request(t *testing.T) {
	t.Parallel()

	t.Run("ToAsn1Der", func(t *testing.T) {
		t.Parallel()

		t.Run("empty", func(t *testing.T) {
			t.Parallel()

			var empty qa.Asn1Request
			encoded, e := empty.ToAsn1Der()
			if nil != e {
				t.Fatalf("unexpected error: %v", e)
			}

			if 0 == len(encoded) {
				t.Fatal("empty bytes got")
			}
		})
	})
}
