package hreq2asn1

import (
	"encoding/asn1"
	"iter"
	"maps"
	"net/http"
	"os"
)

type HttpRequest struct {
	*http.Request
}

func (r HttpRequest) ToMethod() HttpMethod {
	switch r.Method {
	case http.MethodGet:
		return HttpMethodGet
	case http.MethodHead:
		return HttpMethodHead
	case http.MethodPost:
		return HttpMethodPost
	case http.MethodPut:
		return HttpMethodPut
	case http.MethodPatch:
		return HttpMethodPatch
	case http.MethodDelete:
		return HttpMethodDelete
	case http.MethodConnect:
		return HttpMethodConnect
	case http.MethodOptions:
		return HttpMethodOptions
	case http.MethodTrace:
		return HttpMethodTrace
	default:
		return HttpMethodUnspecified
	}
}

func (r HttpRequest) ToHeaders() []HeaderItem {
	var pairs iter.Seq2[string, []string] = maps.All(r.Request.Header)
	var ret []HeaderItem
	for key, values := range pairs {
		for _, val := range values {
			ret = append(ret, HeaderItem{
				Key: key,
				Val: val,
			})
		}
	}
	return ret
}

func (r HttpRequest) ToAsn1Request() Asn1Request {
	return Asn1Request{
		HttpMethod:    r.ToMethod(),
		Url:           r.Request.URL.String(),
		Protocol:      r.Request.Proto,
		ContentLength: r.Request.ContentLength,
		Host:          r.Request.Host,
		RemoteAddr:    r.Request.RemoteAddr,
		Headers:       r.ToHeaders(),
	}
}

type HeaderItem struct {
	Key string
	Val string
}

type HttpMethod asn1.Enumerated

const (
	HttpMethodUnspecified HttpMethod = 0
	HttpMethodGet         HttpMethod = 1
	HttpMethodHead        HttpMethod = 2
	HttpMethodPost        HttpMethod = 3
	HttpMethodPut         HttpMethod = 4
	HttpMethodPatch       HttpMethod = 5
	HttpMethodDelete      HttpMethod = 6
	HttpMethodConnect     HttpMethod = 7
	HttpMethodOptions     HttpMethod = 8
	HttpMethodTrace       HttpMethod = 9
)

type Asn1Request struct {
	HttpMethod
	Url           string `asn1:"utf8"`
	Protocol      string `asn1:"ia5"`
	ContentLength int64
	Host          string `asn1:"utf8"`
	RemoteAddr    string `asn1:"ia5"`
	Headers       []HeaderItem
}

func (a Asn1Request) ToAsn1Der() ([]byte, error) {
	return asn1.Marshal(a)
}

func (a Asn1Request) ToAsn1DerFs(filename string) error {
	f, e := os.Create(filename)
	if nil != e {
		return e
	}
	defer f.Close()

	der, e := a.ToAsn1Der()
	if nil != e {
		return e
	}
	_, e = f.Write(der)
	return e
}
