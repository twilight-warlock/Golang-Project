package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type newsletter struct {
	title  string
	body   string
	topic  string
	author string
}

func main() {

	r := mux.NewRouter()
	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", "gavin:hooli#123@tcp(127.0.0.1:3306)/my_db")
  if err != nil {
    panic(err.Error())
  }
	defer db.Close()

	r.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		// get user email from r
		// add it to list of users subscribed to a topic in db
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var result map[string]interface{}
		json.Unmarshal([]byte(string(body)), &result)
		result2, err := db.Exec("INSERT INTO subscribers (email) VALUES (?)", result["email"])
		fmt.Println(result2, err)
		fmt.Fprintf(w, "Send some response")
	}).Methods("POST")


	r.HandleFunc("/unsubscribe", func(w http.ResponseWriter, r *http.Request) {
		// get user email from r
		// add it to list of users subscribed to a topic in db
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var result map[string]interface{}
		json.Unmarshal([]byte(string(body)), &result)
		result2, err := db.Exec("DELETE FROM subscribers WHERE email=(?)", result["email"])
		fmt.Println(result2, err)
		fmt.Fprintf(w, "Send some response")
	}).Methods("POST")

	r.HandleFunc("/send_newsletter", func(w http.ResponseWriter, r *http.Request) {
		// get user topic,head,content from r
		// Add it to collection topic inside a new document and call sendMail function on it
		rows, err := db.Query("SELECT email FROM subscribers")
    if err != nil {
				fmt.Println("Err")
        return
    }
    defer rows.Close()

    // An album slice to hold data from returned rows.
    var emails []string

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var email string
        if err := rows.Scan(&email); err != nil {
					fmt.Println("Err")
            return
        }
        emails = append(emails, email)
    }
    if err = rows.Err(); err != nil {
				fmt.Println("Err")
        return
    }
		fmt.Println(emails)
		

		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)

		var result map[string]interface{}
		json.Unmarshal([]byte(string(body)), &result)

		fmt.Println("Title :", result["title"])
		fmt.Println("Author :", result["author"])
		fmt.Println(reflect.TypeOf(result["title"]))
		fmt.Println(result)
		subject := result["title"].(string)
		content :=result["body"].(string)
		sendMail(subject, content,emails)
	}).Methods("POST")

	http.Handle("/", r)
	http.ListenAndServe(":8080", r)

}
