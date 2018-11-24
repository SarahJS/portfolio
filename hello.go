package main

import (
	"net/http"
	"time"
	"html/template"
	"log"
)

type PageVariables struct {
	Date 	string
	Time 	string
}

func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	IndexVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, IndexVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8998", nil)
}
