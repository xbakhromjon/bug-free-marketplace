package rest

import (
	"github.com/go-chi/render"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err error) {

	// default handler
	writer.WriteHeader(http.StatusInternalServerError)
	render.JSON(writer, request, err.Error())
}
