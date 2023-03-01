package controller

import (
	"encoding/json"
	"fmt"
	"github.com/database"
	"github.com/gorilla/mux"
	"github.com/model"
	"github.com/repository"
	"github.com/response"
	"net/http"
	"strconv"
)

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := r.URL.Query().Get("user")

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	users, error := repository.NewUserRepository(db).FindAll(nameOrNick)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	user, error := repository.NewUserRepository(db).FindById(id)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	if user.Id == 0 {
		response.Error(w, http.StatusNotFound, fmt.Errorf("user not found with ID [%d]", id))
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	if error = repository.NewUserRepository(db).DeleteById(id); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, error := requestBody(r)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := user.Prepare(); error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	newId, error := repository.NewUserRepository(db).Create(user)
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	user.Id = newId

	response.JSON(w, http.StatusCreated, user)
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	user, error := requestBody(r)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	if error := user.Prepare(); error != nil {
		response.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	parameters := mux.Vars(r)

	id, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		response.Error(w, http.StatusBadRequest, error)
		return
	}

	user.Id = id

	db, error := database.Connect()
	if error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	if error := repository.NewUserRepository(db).Update(user); error != nil {
		response.Error(w, http.StatusInternalServerError, error)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func requestBody(r *http.Request) (model.User, error) {
	requestBody, error := requestBodyToString(r)
	if error != nil {
		return model.User{}, error
	}

	var user model.User

	if error = json.Unmarshal(requestBody, &user); error != nil {
		return model.User{}, error
	}

	return user, nil
}
