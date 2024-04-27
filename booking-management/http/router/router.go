package router

import (
	"booking-management/http/transport/response"
	"booking-management/utils/container"
	"net/http"

	"github.com/gorilla/mux"
)

func Init(ctr container.Containers) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response.Send(w, []byte("Notes APP"), http.StatusOK)
	}).Methods(http.MethodGet)

	return r
}
