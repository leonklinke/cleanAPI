package infrastructure

import (
	"cleanApi/interfaces"
	"cleanApi/usecases"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func ServeRoutes(logger usecases.Logger) {
	router := mux.NewRouter()
	router = CreatingRoutes(router, logger)

	fmt.Println("serving")

	err := http.ListenAndServe("0.0.0.0:"+os.Getenv("API_PORT"), router)
	if err != nil {
		logger.LogError("%s", err)
	}
}

func CreatingRoutes(router *mux.Router, logger usecases.Logger) *mux.Router {
	databaseHandler, err := NewDatabaseHandler()
	if err != nil {
		logger.LogError("%s", err)
	}

	userController := interfaces.NewUserController(databaseHandler, logger)

	//User
	router.HandleFunc("/user", userController.Show).
		Methods("GET")

	return router
}
