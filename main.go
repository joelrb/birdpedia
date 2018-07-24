// https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func main() {

	fmt.Println("Starting server...")
	connString := "dbname=bird_encyclopedia sslmode=disable"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()

	//pass router
	http.ListenAndServe(":8080", r)
}

func newRouter() *mux.Router {
	//declare a new router
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Static file declaration
	staticFileDirectory := http.Dir("./assets/")

	// Asset handler
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	// Path prefix
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")

	// Bird handler router
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")

	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
