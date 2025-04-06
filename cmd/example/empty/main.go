package main

import (
	"net/http"
	"os"

	qa "github.com/takanoriyanagitani/go-http-req2asn1"
)

func Must[T any](t T, e error) T {
	if nil != e {
		panic(e)
	}
	return t
}

var method string = http.MethodGet

var hreqp *http.Request = Must(http.NewRequest(
	method,
	"http://example.com/",
	nil,
))

var hreq qa.HttpRequest = qa.HttpRequest{Request: hreqp}

var areq qa.Asn1Request = hreq.ToAsn1Request()

func main() {
	der, e := areq.ToAsn1Der()
	if nil != e {
		panic(e)
	}

	_, e = os.Stdout.Write(der)
	if nil != e {
		panic(e)
	}
}
