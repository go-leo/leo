package util

import "net/http"

func GetContentType(header http.Header) string {
	return header.Get("Content-Type")
}

func GetAcceptType(header http.Header) string {
	return header.Get("Accept")
}

func SetContentType(header http.Header, v string) {
	header.Set("Content-Type", v)
}

func SetAcceptType(header http.Header, v string) {
	header.Set("Accept", v)
}
