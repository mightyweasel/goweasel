package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)


func IndexHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("IndexHandler call" + rs)
	status := "weasel index"

	wv := WebView {
		Title: status,
	}
	// Render page
	RenderRoute(w, wv)
}
func GuideHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("GuideHandler call: " + rs)
	status := "weasel guide"

	wv := WebView {
		Title: status,
	}
	// Render page
	RenderRoute(w, wv)
	//Render(w, "weaseltemplates/200.html", wv)
}
func SearchHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("SearchHandler call" + rs)
	status := "do wit.ai: "

	// http://localhost:8080/search?recognition_language=en-US&recognition_lang_tlxd=en-US&weasel_ask=why+should+my+boss+care+about+open+source
	fmt.Println("GET params were:", req.URL.Query())

	recognitionlanguage := req.FormValue("recognition_language")
	autotranslate := req.FormValue("recognition_lang_tlxd")
	searchquery := req.FormValue("weasel_ask")

	vars := mux.Vars(req)
	q := vars["q"]
	status = status + " " + q + recognitionlanguage + " " + autotranslate + " " + searchquery 

	// todo: next step is to update the webview, match it to the objects we expect
	wv := WebView {
		Title: status,
	}
	// Render page
	//Render(w, "weaseltemplates/gc-ermine.html", wv)
	RenderRoute(w, wv)
}
func RedirectHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	rs=""
	fmt.Println("RedirectHandler call" + rs)
	vars := mux.Vars(req)
	q := vars["q"]
	
	log.Println("Redirecting to home: " + q)
	http.Redirect(w, req, "/", 302)
	return
}

func WeaselCreateRouter(port string) {
	decoder  := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	r := mux.NewRouter()	
   
    r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/guide/{q}", GuideHandler)
	r.HandleFunc("/search", SearchHandler)
	r.HandleFunc("/api", RedirectHandler)	
	
	// Note: This is the moment I realized how cool Go actually is.
	// It's been fighting me until now to get a basic web page to render.
	// I just wanted it to get out of my may and do what I expected, but that the thing
	// Go isn't like the other ones. It's a completely different mindset.
	// You must be flexible like the branches of a tree lest ye break in the cold canadian winters
	// 
	// Also Sin's Cat/Dog Hypothesis applies here: If you pet a cat like you would a dog, you're gonna
	// get bitten in all probability. If you pet a dog like a cat, well thats an exercise for the reader
	// most large breeds of dog are heavier than I am. It's still a jungle, no matter how much AI
	// we endeavor to build
	//
	// Enough philosophy, let's get back to the coolness of Go
	//
	// so Go is a bit different, your files dont exist unless you say they do.
	// we're going to invoke magic to create directory that doesnt really exist root/css/
	// that we can use in our templates (so all css can live together effectively)
	// same for js. You can use this in a lot more powerful ways. These are just the basics
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("weaselstatic/css/"))))

	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
		http.FileServer(http.Dir("weaselstatic/js/"))))

	fmt.Println("Starting webserver on port " + port)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+port, r))
}