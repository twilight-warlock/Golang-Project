package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		// get user email from r
		// add it to list of users subscribed to a topic in db
		fmt.Fprintf(w, "Send some response")
	})
	r.HandleFunc("/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		// get user email from r
		// remove it from list of users subscribed to a topic in db
		fmt.Fprintf(w, "Send some response")
	})

	r.HandleFunc("/createNewsletter", func(w http.ResponseWriter, r *http.Request) {
		// get user topic,head,content from r
		// Add it to collection topic inside a new document and call sendMail function on it
		sendMail()
		fmt.Fprintf(w, "Send some response")
	})


	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
	fmt.Println("Listening on 8080")

}
