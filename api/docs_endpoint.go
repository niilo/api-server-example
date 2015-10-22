package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

func createDocsRoute(e *echo.Echo) {
	for apiKey, _ := range apiDescriptionsJson {
		e.Get("/docs/"+apiKey, apiDescriptionHandler)
	}
}

func apiDescriptionHandler(w http.ResponseWriter, r *http.Request) {
	apiKey := strings.TrimLeft(r.RequestURI, "docs/")

	if json, ok := apiDescriptionsJson[apiKey]; ok {
		w.Write([]byte(json))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
