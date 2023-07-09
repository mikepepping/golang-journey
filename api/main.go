package main

import (
	"os"
  "fmt"
  "path/filepath"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/me", asContentType("application/json", MeHandler))
	http.ListenAndServe(":8080", router)
}


// BEGIN DataStore
type DataStore struct {
  RootDir string
  FileName string
}

func (ds DataStore) InitDataStore() {
  _, err := os.Stat(ds.RootDir)
  if (os.IsNotExist(err)) {
    fmt.Println("ERROR - data store directory does not exist")
    panic(err)
  }
}


func (ds DataStore) Store(dataType string, id string, data any) {
  filename := id + ".json"
  fullFilepath := filepath.Join(ds.RootDir, dataType, filename)
  file, err := os.Create(fullFilepath)
  if (err != nil) {
    fmt.Printf("ERROR - Could not store %s with id %s in file %s", dataType, id, fullFilepath)
    panic(err)
  }

  file.Close()
}

// END DataStore

type UserRow struct {
  Email string
  GivenName string
  LastName string
  Password string
}

func (ur UserRow) Id() string {
  return ur.Email
}

func (ur UserRow) DataType() string {
  return "user"
}


// "middleware" for setting a routes content type
func asContentType(contentType string, next http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", contentType)

    next.ServeHTTP(w, r)
  })
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	serialized, err	:= json.Marshal(struct{}{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(serialized))
}

