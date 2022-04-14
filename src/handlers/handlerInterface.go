package handlers

import "net/http"

type HandlerI interface {
	Get(rw http.ResponseWriter, rq *http.Request)
	Put(rw http.ResponseWriter, rq *http.Request)
	Post(rw http.ResponseWriter, rq *http.Request)
}
