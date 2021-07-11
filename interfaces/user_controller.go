package interfaces

import (
	"cleanApi/usecases"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UserController struct {
	UserInteractor usecases.UserInteractor
	Logger         usecases.Logger
}

func NewUserController(databaseHandler DatabaseHandler, logger usecases.Logger) *UserController {
	return &UserController{
		UserInteractor: usecases.UserInteractor{
			UserRepository: &UserRepository{
				DatabaseHandler: databaseHandler,
			},
		},
		Logger: logger,
	}
}

func (controller *UserController) Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside user controller")
	controller.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	userId, _ := strconv.Atoi(r.URL.Query().Get("id"))

	user, err := controller.UserInteractor.Show(userId)
	if err != nil {
		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(user)
}
