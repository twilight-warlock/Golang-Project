package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type newsletter struct{
	title string
	body string
	topic string
	author string
}

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

	r.HandleFunc("/create_newsletter", func(w http.ResponseWriter, r *http.Request) {
		// get user topic,head,content from r
		// Add it to collection topic inside a new document and call sendMail function on it

		
		// err := json.NewDecoder(r.Body).Decode(&newsletter)

		// fmt.Println(err)
		// if err != nil {
		// 		// Simplified
		// 		fmt.Println(err)
		// 		return
		// }
		// fmt.Println(newsletter)
		
    // post, _ := range r.Body.Read(){
		// 	fmt.Println(post)
		// }

		defer r.Body.Close()
    body, _ := ioutil.ReadAll(r.Body)

    // ioutil.WriteFile("dump", body, 0600)
		var iot newsletter
    err := json.Unmarshal(body, &iot)
		fmt.Println(err)
    fmt.Println("done")
		fmt.Println(iot)


		// sendMail()
	})


	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
	

}
