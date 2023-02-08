package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()
	return router
}
