package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// landingHandler handles requests to "/"
func landingHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("landing.html")
	t.Execute(w, nil)
}

// landingHandler handles requests of the form "/{username}".
func achievementHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	userID, err := getUserID(username)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	unearned, err := unearnedAchievements(userID)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}
	categorized := categorizeAchievements(unearned)
	t, _ := template.ParseFiles("achievements.html")
	t.Execute(w, categorized)
}