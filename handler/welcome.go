package handler

import (
	"encoding/json"
	"fmt"
	"github.com/tejashwikalptaru/tutorial/database/helper"
	"github.com/tejashwikalptaru/tutorial/middlewares"
	"github.com/tejashwikalptaru/tutorial/models"
	"net/http"
)

func isErr(err error, typeErr string) bool {
	return err.Error() == "pq: duplicate key value violates unique constraint "+typeErr
}

// Greet function

//func Greet(writer http.ResponseWriter, request *http.Request) {
//	var userDetail models.User
//	err := json.NewDecoder(request.Body).Decode(&userDetail)
//	fmt.Println("Hai->>")
//	userID, err := helper.CreateUser(userDetail.Name, userDetail.Email)
//	if isErr(err, models.ActiveUser) {
//
//	}
//	if err != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//
//	user, userErr := helper.GetUser(userID)
//	if userErr != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	jsonData, jsonErr := json.Marshal(user)
//	if jsonErr != nil {
//		writer.WriteHeader(http.StatusInternalServerError)
//		return
//	}
//	writer.Write(jsonData)
//}

func Login(writer http.ResponseWriter, request *http.Request) {
	//	Login the Registered Use

	var logIN models.Login
	err := json.NewDecoder(request.Body).Decode(&logIN)

	if err != nil {
		writer.Write([]byte("Some thing Went Gone"))
		writer.WriteHeader(http.StatusBadRequest)
	}

	present, userId := helper.LoginDetails(logIN.Email, logIN.Password)
	if !present {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("Invalid Id Password"))
		return
	}

	err, token := helper.SessionLogin(userId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("Some Thing Went Wrong"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	outboundMsg := map[string]interface{}{
		"token": token,
	}

	// Send toke to client
	err = json.NewEncoder(writer).Encode(outboundMsg)

	if err != nil {
		writer.WriteHeader(400)
		writer.Write([]byte("Session Token is not Generated"))
	}

}

func Logout(writer http.ResponseWriter, request *http.Request) {
	// Logout the User or update sessions table -> change logout time from null to timeNow
	userDetail := middlewares.UserContext(request)
	err := helper.SetLogout(userDetail)
	if err != nil {
		writer.Write([]byte("Some thing went Wrong"))
		return
	}

}

func Register(writer http.ResponseWriter, request *http.Request) {
	// Add/Register new user in the user table

	var userDetail models.User
	err := json.NewDecoder(request.Body).Decode(&userDetail)
	writer.Write([]byte(userDetail.Name))
	//fmt.Println("Hai->>")
	str, err := helper.CreateUser(userDetail.Name, userDetail.Email, userDetail.Password)
	if str == " Data Already Exist" {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(str))
		return
	}
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func AddTask(writer http.ResponseWriter, request *http.Request) {
	//	update/add more task in task table

	var addTask models.Task

	err := json.NewDecoder(request.Body).Decode(&addTask)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("Some Thing Went Wrong"))
		return
	}

	userDetail := middlewares.UserContext(request)
	taskId, err := helper.CreateTask(addTask, *userDetail)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Please Try one more Time!"))
		return
	}
	fmt.Println(taskId)

}

func UpdateTask(writer http.ResponseWriter, request *http.Request) {
	//	change the status of the task or change the deadline of the task

}

func GetTask(writer http.ResponseWriter, request *http.Request) {
	//	Get all the active task in the task table

	userDetail := middlewares.UserContext(request)

	list := helper.GetTask(userDetail)
	err := json.NewEncoder(writer).Encode(list)
	if err != nil {
		writer.WriteHeader(500)
		return
	}

}
