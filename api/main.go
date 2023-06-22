package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type User struct {
	Id 	string `json:"id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/me", MeHandler)
	http.ListenAndServe(":8080", router)
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	user := User{ "id", "name", "email@email.com" }
	serialized, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(serialized))
	}
	w.Write([]byte(serialized))
}

