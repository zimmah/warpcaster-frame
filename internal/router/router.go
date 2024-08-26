package router

import "net/http"

func Router() {
	mux := http.NewServeMux()
	mux.Handle("GET /api/frames/{frameID}", logger(http.HandlerFunc(handleGetFrameByID)))
}