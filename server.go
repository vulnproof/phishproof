package main

import (
	"html/template"
	"log"
	"net/http"
)

func getFormData() string {
	urlToCheck := ""
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, err := template.ParseFiles("html/form.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else if r.Method == http.MethodPost {
			urlToCheck = r.FormValue("url")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return urlToCheck
}

func servTheApp() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
