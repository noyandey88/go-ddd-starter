package middlewares

import "net/http"

type Middlware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlwares []Middlware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlwares: make([]Middlware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middlware) {
	mngr.globalMiddlwares = append(mngr.globalMiddlwares, middlewares...)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middlware) http.Handler {

	h := handler

	for _, middleware := range middlewares {
		h = middleware(h)
	}

	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler

	for _, middleware := range mngr.globalMiddlwares {
		h = middleware(h)
	}

	return h
}
