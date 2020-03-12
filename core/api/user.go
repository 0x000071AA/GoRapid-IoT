package api

import (
	"encoding/json"
	"net/http"
	"src/GoRapid-IoT/core/databases"
)

var sqlClient = databases.SqlConnect("")

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	err := sqlClient.Query("", "")
	if err != nil {
		e := HTTPInternalServerError(err.Error())
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}

	if true {
		e := HTTPNotFound("User not found")
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}
	res := HTTPGet("")
	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if r.Body == nil {
		e := HTTPBadRequest("No body provided")
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		e := HTTPBadRequest("Invalid json")
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}

	user, errInsert := sqlClient.Insert("")
	if errInsert != nil {
		e := HTTPInternalServerError(errInsert.Error())
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}
	res := HTTPCreated("user created")
	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	err := sqlClient.Delete("")
	if err != nil {
		e := HTTPInternalServerError(err.Error())
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}

	res := HTTPDeleted("user deleted")
	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if r.Body == nil {
		e := HTTPBadRequest("No body provided")
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		e := HTTPBadRequest("Invalid json")
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}

	_, errUpdate := sqlClient.Update("")
	if errUpdate != nil {
		e := HTTPInternalServerError(errUpdate.Error())
		w.WriteHeader(e.status)
		json.NewEncoder(w).Encode(e)
		return
	}
	res := HTTPUpdated("user updated")
	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(res)
}
