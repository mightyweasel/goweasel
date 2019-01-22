package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"html/template"
	"strings"
	
	//"github.com/mightyweasel/goweasel/weaselroutehandler"	
	//"github.com/mightyweasel/goweasel/weaselgui"

	"github.com/gorilla/mux"
	//"github.com/gorilla/schema"
	"github.com/joho/godotenv"

	//"github.com/go-pg/pg"
)


//const fav_icon_url = "http://www.canada.ca/etc/designs/canada/wet-boew/assets/favicon.ico"


// WebView is a framework to send objects & data to a Web view
type WebView struct {
	Redirect    bool
	Title string
	CategoryMap map[string]int
	Counter    []int
	Flashes []interface{}
}

// SplitLines transfomrs results text string into slice
func SplitLines(s string) []string {
	sli := strings.Split(s, "/n")
	return sli
}

func sliceString(s string, i int) string {

	l := len(s)

	if l > i {
		return s[:i] + "..."
	}
	return s[:l]
}

func subtract(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func percent(a, b int) float32 {
	return (float32(a) / float32(b)) * 100.0
}

func isIn(s []int, t int) bool {
	for _, n := range s {
		if n == t {
			return true
		}
	}
	return false
}

func isInString(s []string, t string) bool {
	for _, n := range s {
		if n == t {
			return true
		}
	}
	return false
}

// Render combines templates, funcs and renders all Web pages in the app
func RenderRoute(w http.ResponseWriter, filename string, data interface{}) {

	//baseTemplate := "weaseltemplates/gc-ermine.html"
	//baseTemplateRef := "base"
	baseTemplate := "weaseltemplates/user-status.html"
	baseTemplateRef := "tmpluserstatus"

    t, _ := template.ParseFiles( baseTemplate ) //setp 1
	t.ExecuteTemplate(w, baseTemplateRef, data)
    //t.Execute(w, "Hello World!") //step 2
}


// AboutHandler renders a character in a Web page
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	fmt.Println("IndexHandler call" + rs)
	status := "weasel index"

	wv := WebView {
		Title: "Click counter: " + status,
	}
	// Render page
	RenderRoute(w, "weaseltemplates/gc-ermine.html", wv)
}

// AboutHandler renders a character in a Web page
func GuideHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	fmt.Println("GuideHandler call: " + rs)
	status := "weasel guide"

	wv := WebView {
		Title: "Click counter: " + status,
	}
	// Render page
	RenderRoute(w, "weaseltemplates/gc-ermine.html", wv)
	//Render(w, "weaseltemplates/200.html", wv)
}


// AboutHandler renders a character in a Web page
func SearchHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	fmt.Println("SearchHandler call" + rs)
	status := "weasel search"
	vars := mux.Vars(req)
	q := vars["q"]
	q = q + " " + status

	wv := WebView {
		Title: "Click counter: " + status,
	}
	// Render page
	//Render(w, "weaseltemplates/gc-ermine.html", wv)
	RenderRoute(w, "weaseltemplates/gc-ermine.html", wv)
}

func RedirectHandler(w http.ResponseWriter, req *http.Request) {
	rs := fmt.Sprintf("%#v", req)
	fmt.Println("RedirectHandler call" + rs)
	vars := mux.Vars(req)
	q := vars["q"]
	
	log.Println("Redirecting to home: " + q)
	http.Redirect(w, req, "/", 302)
	return
}


func main() {
	env := os.Getenv("ENVIRONMENT")
	if "" == env {
	  env = "development"
	}
	godotenv.Load(".env." + env + ".local")

	if os.Getenv("ENVIRONMENT") == "production" {
		fmt.Println("Init "+os.Getenv("ENVIRONMENT")+" Webserver")
	}
	if os.Getenv("ENVIRONMENT") == "development" {
		fmt.Println("Init "+os.Getenv("ENVIRONMENT")+" Webserver")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	port = "8080"

	//decoder  = schema.NewDecoder()
	//decoder.IgnoreUnknownKeys(true)

	r := mux.NewRouter()

	fmt.Println("Starting "+os.Getenv("ENVIRONMENT")+" Webserver at port " + port)
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/guide/", GuideHandler)
	r.HandleFunc("/search/{q}", SearchHandler)
	r.HandleFunc("/redirect/{q}", RedirectHandler)
	
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+port, r))
}

