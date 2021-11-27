package api

import (
	"encoding/json"
	"fmt"
	"forum/internal/app/models"
	"forum/pkg/logger"
	"net/http"
)

// Message Вспомогательная структура для формирования сообщений
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHandlers(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	initHandlers(w)
	logger.InfoLogger.Println("Post User Register POST /api/user/register ")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		logger.InfoLogger.Println("Invalid json received from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}

	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		logger.InfoLogger.Println("Troubles while accessing database table (users) with id. err:", err)
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles to accessing database. Try again",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}

	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("User {login:%s} successfully registered!", userAdded.Login),
		IsError:    false,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(msg)

}

func (api *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	initHandlers(w)

	users, err := api.storage.User().SelectAll()
	if err != nil {
		logger.InfoLogger.Println(err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some troubles to assessing users in database. Try later",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	logger.InfoLogger.Println("Ger All Users GET /users")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
