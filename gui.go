package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"io/ioutil"
	"html/template"
)


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

var templates *template.Template
func InitalizeGuiTemplates() {

    var allFiles []string
    files, err := ioutil.ReadDir("./weaseltemplates")
    if err != nil {
        fmt.Println(err)
    }
    for _, file := range files {
        filename := file.Name()
        if strings.HasSuffix(filename, ".tmpl") {
            allFiles = append(allFiles, "./weaseltemplates/"+filename)
        }
    }

	templates, err = template.ParseFiles(allFiles...) //parses all .tmpl files in the 'templates' folder

}

func RenderRoute(w http.ResponseWriter, data interface{}) {
	t := templates.Lookup("base.html.tmpl")
    //s1.ExecuteTemplate(os.Stdout, "header", nil)
    //fmt.Println()
    
	// this line blow up with escaped html
	if err := t.ExecuteTemplate(w, "base", data); err != nil {
    // this line is fine, it routes, but renders nuffin
    //if err := t.Execute(w, nil); err != nil {
        log.Printf("Failed to execute template: %v", err)
    }	
}
