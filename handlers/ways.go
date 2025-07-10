package handlers

import (
	"net/http"
)

func RouterWays() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /{BucketName}", CreateBucket)
	mux.HandleFunc("GET /", ListBuckets)
	mux.HandleFunc("DELETE /{BucketName}", DeleteBucket)

	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", CreateObject)
	mux.HandleFunc("GET /{BucketName}", ListObjects)
	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", GetObject)
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", DeleteObject)

	return mux
}
