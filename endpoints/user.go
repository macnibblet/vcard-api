package endpoints

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/macnibblet/vcard-api/cards"
	"github.com/macnibblet/vcard-api/database"
	"net/http"
)

type UserEndpoints struct {
	repo *database.UserRepository
}

func (e *UserEndpoints) create(resp http.ResponseWriter, req *http.Request) {
	email := req.Form.Get("email")

	// User already exists
	if _, err := e.repo.GetByEmail(email); err == nil {
		resp.WriteHeader(http.StatusConflict)
		return
	}

	user := &database.User{
		Email:    email,
		Password: req.Form.Get("password"),
	}

	if err := e.repo.Create(user); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write(RESP_INTERNAL_SERVER_ERROR)
		return
	}

}

func (e *UserEndpoints) getById(resp http.ResponseWriter, req *http.Request) {

}

func InjectUserRoutes(context *cards.Context, router *mux.Router) {
	api := UserEndpoints{
		repo: context.UserRepository,
	}

	router.HandleFunc("/users", api.create).Methods(http.MethodPost)
	router.Handle("/users/{uuid}", negroni.New(
		negroni.HandlerFunc(context.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(wrapHandlerFunc(api.getById)),
	)).Methods(http.MethodGet)
}
