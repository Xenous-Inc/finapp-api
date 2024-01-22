package responser

import (
	"net/http"

	"github.com/Xenous-Inc/finapp-api/internal/dto"
	"github.com/go-chi/render"
)

func BadRequset(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusBadRequest)
	render.JSON(w, r, &dto.ApiError{
		Error: msg,
	})
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusUnauthorized)
	render.JSON(w, r, &dto.ApiError{
		Error: "Unauthorized",
	})
}

func Internal(w http.ResponseWriter, r *http.Request, msg string) {
	render.Status(r, http.StatusInternalServerError)
	render.JSON(w, r, &dto.ApiError{
		Error: msg,
	})
}

func Success(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, data)
}
