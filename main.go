package main

import (
	"github.com/gabielfemi/my-version-social-tracker/tracker"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", Index)
	router.HandleFunc("/search", Search).Methods("GET")
	router.PathPrefix("/assets").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))


	server := &http.Server{
		Addr: "127.0.0.1:8000",
		Handler: router,

	}
	log.Println("Listening on 127.0.0.1:8000")
	log.Fatalln(server.ListenAndServe())

}

func Index(w http.ResponseWriter, r *http.Request) {
	tracker.Render(w, "index.html", r)
}

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//_, _ = fmt.Fprintf(w, "You searched for %s", r.FormValue("searchPhrase"))
		tracker.Render(w, "search.html", r)
	} else {
		// http.Redirect(w, r , "/", 301)
	}

}