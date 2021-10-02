package controller

import (
	"academy-go-q32021/infrastructure/datastore"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *datastore.Db
	Router *mux.Router
}

func InitializeApp(db *datastore.Db) *App {
	app := new(App)
	app.DB = db
	return app
}

func (app *App) Run(router *mux.Router) {
	app.Router = router

	http.ListenAndServe(":8000", router)
}

type RequestHandlerFunction func(db *datastore.Db, w http.ResponseWriter, r *http.Request)

func (app *App) HandleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(app.DB, w, r)
	}
}
