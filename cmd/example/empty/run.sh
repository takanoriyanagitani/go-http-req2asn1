#!/bin/sh

outname=./sample.http.request.asn1.der.dat

echo serializing an http request...
./empty |
	dd \
		if=/dev/stdin \
		of="${outname}" \
		bs=1048576 \
		status=none

echo
echo "decoding/showing the DER bytes using python's asn1tools and jq..."
cat "${outname}" |
	python3 \
		-m uv \
		run \
		./sample.py |
    jq
