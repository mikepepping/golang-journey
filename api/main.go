package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Email string `json:"email"`
}

func (sf *SignupForm) isValid() bool {
	emailValid := len(sf.Email) > 0
	passwordValid := len(sf.Password) > 0

	return emailValid && passwordValid
}

type UserResponse struct {
	Email string `json:"email"`
}

type AccessReponse struct {
	AccessToken string `json:"accessToken"`
}

type ErrorResponse struct {
	Code    uint   `json:"code"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

var db *sql.DB
var err error

func main() {
	dbConnectionString := os.Getenv("AUTH_DB_CONNECTION_STRING")
	if len(dbConnectionString) <= 0 {
		fmt.Println("AUTH_DB_CONNECTION_STRING is not set")
		return
	}

	db, err = sql.Open("sqlite3", dbConnectionString)
	if err != nil {
		fmt.Println("Failed to open database")
		return
	}
	defer db.Close()

	err = InitDb(db)
	if err != nil {
		fmt.Println("FAILED to initialize Database")
		return
	}
	fmt.Println("Initialized Database")

	router := mux.NewRouter()
	router.HandleFunc("/signup", asContentType("application/json", SignupHandler)).Methods("POST")
	router.HandleFunc("/me", asContentType("application/json", MeHandler)).Methods("GET")
	http.ListenAndServe(":8080", router)

	fmt.Println("Shutting down")
}

func InitDb(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (email varchar(255), password varchar(255))")
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// "middleware" for setting a routes content type
func asContentType(contentType string, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", contentType)

		next.ServeHTTP(w, r)
	})
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var signupForm SignupForm
	err := json.NewDecoder(r.Body).Decode(&signupForm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if signupForm.isValid() != true {
		http.Error(w, "Users details were invalid", http.StatusUnprocessableEntity)
		return
	}

	existsStmt, err := db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		http.Error(w, "Failed to check if user exists", http.StatusInternalServerError)
		return
	}
	defer existsStmt.Close()

	rows, err := existsStmt.Query(signupForm.Email)
	if err != nil {
		http.Error(w, "Failed to check if user exists", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows.Next() {
		http.Error(w, "A user with this email already exists", http.StatusConflict)
		return
	}

	insertStmt, err := db.Prepare("INSERT INTO users (email, password) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare new user", http.StatusInternalServerError)
		return
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(signupForm.Email, signupForm.Password)
	if err != nil {
		http.Error(w, "Failed to create new user", http.StatusInternalServerError)
		return
	}

	response := SignupResponse{signupForm.Email}
	serialized, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(serialized))
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	serialized, err := json.Marshal(struct{}{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(serialized))
}
