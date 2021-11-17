package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  r := mux.NewRouter();

	// send mail 
	// subscribe to newsletter
	// unsubscribe from newsletter
	// create new newsletter
	// delete newsletter

  //Regestering routes

	r.HandleFunc("/sendMail", func(w http.ResponseWriter, r *http.Request){
		// get all users email ids from database
		// call the send method from mail service to send all recipients the newsletter
    fmt.Fprintf(w, "Sending Mail");
  })

	r.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request){
		// get user email from r
		// send it to list of users subscribed to a topic in db
    fmt.Fprintf(w, "Send some response");
  })

	r.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request){
		// get user email from r
		// send it to list of users subscribed to a topic in db
    fmt.Fprintf(w, "Send some response");
  })

	r.HandleFunc("/createNewsletter", func(w http.ResponseWriter, r *http.Request){
		// get user topic,head,content from r
		// Add it to collection topic inside a new document and call sendMail function on it
    fmt.Fprintf(w, "Send some response");
  })

	r.HandleFunc("/createNewsletter", func(w http.ResponseWriter, r *http.Request){
		// get user topic,head from r
		// Remove it from collection topic 
    fmt.Fprintf(w, "Send some response");
  })
  
  http.Handle("/", r)
  http.ListenAndServe(":8080", r)

}