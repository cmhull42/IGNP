package routes

import (
	"net/http"

	sysmodel "github.com/cmhull42/ignp/model/system"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// ResourceRoutes returns all the resource routes
func ResourceRoutes(e Env) http.Handler {
	router := chi.NewRouter()

	router.Get("/{resourceID}", e.GetResource)
	router.Get("/", e.GetResources)

	return router
}

// GetResource returns one single resource identified by the resourceid param
func (e Env) GetResource(w http.ResponseWriter, r *http.Request) {
	resource := []sysmodel.Resource{}
	err := e.db.Select(&resource, "select * from SystemResources where id=?", chi.URLParam(r, "resourceID"))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if len(resource) == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	w.WriteHeader(200)
	render.JSON(w, r, resource[0])
}

// GetResources returns all resources
func (e Env) GetResources(w http.ResponseWriter, r *http.Request) {
	resource := []sysmodel.Resource{}
	err := e.db.Select(&resource, "select * from SystemResources")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	if len(resource) == 0 {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	w.WriteHeader(200)
	render.JSON(w, r, resource)
}
