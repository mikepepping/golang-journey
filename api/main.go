package main

import (
	"os"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	propelauth "github.com/propelauth/propelauth-go/pkg"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/me", MeHandler)
	http.ListenAndServe(":8080", router)
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	propelApiKey := os.Getenv("PROPELAUTH_API_KEY")
	propelAuthUrl := os.Getenv("PROPELAUTH_AUTH_URL")
	propel_client, err := propelauth.InitBaseAuth(propelAuthUrl, propelApiKey, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := propel_client.GetUser(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)

	serialized, err	:= json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(serialized))
}

