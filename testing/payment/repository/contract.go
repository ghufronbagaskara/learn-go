package repository

import "net/http"

//go:generate mockgen -build_flags=--mod=mod -destination=mock/xendit_mock.go -package=mock . HttpConector
type HttpConector interface {
	Do(req *http.Request) (*http.Response, error)
}