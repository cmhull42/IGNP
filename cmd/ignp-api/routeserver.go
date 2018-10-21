package main

import "net/http"

type RouteServer interface {
	Routes() http.Handler
}
