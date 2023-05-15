package main

import (
	"net/http"
)

// this function is routing to home
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// this function is routing to pageone
func pageOne(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "pageOne.html")
}

// this function is routing to pageone
func aboutUs(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
