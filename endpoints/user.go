package endpoints

import (
	"gopkg.in/pg.v5"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/macnibblet/vcard-api/database"
)

type UserEndpoints struct {
	repo *database.UserRepository
}

func (e *UserEndpoints) create(resp http.ResponseWriter, req *http.Request) {


}

func (e *UserEndpoints) getById(resp http.ResponseWriter, req *http.Request) {

}

func InjectUserRoutes(db *pg.DB, router *mux.Router) {
	api := UserEndpoints{db}

	router.HandleFunc("/users", api.create).Methods(http.MethodPost)
	router.HandleFunc("/users/{uuid}", api.getById).Methods(http.MethodGet)
}
