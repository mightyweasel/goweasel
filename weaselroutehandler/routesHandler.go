package weaselroutehandler

import (
	"fmt"
	"log"
	"net/http"
	"github.com/mightyweasel/goweasel/weaselgui"
)

// AboutHandler renders a character in a Web page
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("IndexHandler call")
	status := "weasel index"

	wv := WebView {
		Title: "Click counter: " + status,
	}
	// Render page
	Render(w, "templates/gc-ermine.html", wv)
}

// AboutHandler renders a character in a Web page
func GuideHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GuideHandler call")
	status := "weasel guide"

	wv := WebView {
		Title: "Click counter: " + status,
	}
	// Render page
	Render(w, "templates/gc-ermine.html", wv)
}


// AboutHandler renders a character in a Web page
func SearchHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("SearchHandler call")
	status := "weasel search"
	vars := mux.Vars(request)
	q := vars["q"]
	q = q + " " + status

	wv := WebView {
		Title: "Click counter: " + status,
	}
	// Render page
	Render(w, "templates/gc-ermine.html", wv)
}

func RedirectHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("RedirectHandler call")
	vars := mux.Vars(request)
	q := vars["q"]
	
	log.Println("Redirecting to home: " + q)
	http.Redirect(w, req, "/", 302)
	return
}
